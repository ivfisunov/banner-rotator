package sqlstorage

import (
	"context"
	"errors"
	"fmt"

	"github.com/ivfisunov/banner-rotator/internal/ucb"
	"github.com/jackc/pgx/v4"
)

type Storage struct {
	dsn string
	db  *pgx.Conn
}

type dbResponse struct {
	BannerID      int
	SlotID        int
	GroupID       int
	Banners       []int
	Groups        []int
	TotalDisplays int
}

func New(dsn string) (*Storage, error) {
	// setup postgres connection
	db, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	return &Storage{dsn: dsn, db: db}, nil
}

func (s *Storage) Close() error {
	err := s.db.Close(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) AddBanner(bannerID, slotID int) error {
	res := dbResponse{}
	// check if banner and slot exist
	err := s.db.QueryRow(
		context.Background(),
		`SELECT
           banners.id AS "bannerId",
           slots.id AS "slotId",
           slots.banners AS "banners"
         FROM banners
         JOIN slots ON slots.id = $2
         WHERE banners.id = $1`,
		bannerID, slotID).Scan(&res.BannerID, &res.SlotID, &res.Banners)
	if err != nil {
		return fmt.Errorf("query error: %w", err)
	}

	for _, v := range res.Banners {
		if v == bannerID {
			return errors.New("banner already has been added")
		}
	}

	// get all social groups
	err = s.db.QueryRow(context.Background(), `SELECT array_agg(id) FROM groups`).Scan(&res.Groups)
	if err != nil {
		return fmt.Errorf("query groups error: %w", err)
	}

	// update slots table
	_, err = s.db.Exec(context.Background(), `UPDATE slots
	    SET banners = array_append(banners, $1)
        WHERE id = $2`,
		res.BannerID, res.SlotID)
	if err != nil {
		return fmt.Errorf("query insert banner error: %w", err)
	}

	// insert new banner statistics
	for _, group := range res.Groups {
		_, err := s.db.Exec(context.Background(), `INSERT INTO stats
        (slot_id, banner_id, group_id)
	    VALUES ($1, $2, $3)`,
			res.SlotID, res.BannerID, group)
		if err != nil {
			return fmt.Errorf("query insert new stats error: %w", err)
		}
	}

	return nil
}

func (s *Storage) DeleteBanner(bannerID, slotID int) error {
	res := dbResponse{}
	// check if banner and slot exist
	err := s.db.QueryRow(
		context.Background(),
		`SELECT banners.id AS "bannerId", slots.id AS "slotId", slots.banners AS "banners"
         FROM banners JOIN slots ON slots.id = $2 WHERE banners.id = $1`,
		bannerID, slotID).Scan(&res.BannerID, &res.SlotID, &res.Banners)
	if err != nil {
		return fmt.Errorf("query error: %w", err)
	}

	_, err = s.db.Exec(context.Background(), `UPDATE slots
	    SET banners = array_remove(banners, $1)
        WHERE id = $2`,
		res.BannerID, res.SlotID)
	if err != nil {
		return fmt.Errorf("query deletion banner error: %w", err)
	}

	// get all social groups
	err = s.db.QueryRow(context.Background(), `SELECT array_agg(id) FROM groups`).Scan(&res.Groups)
	if err != nil {
		return fmt.Errorf("query groups error: %w", err)
	}

	// delete banner from statistics
	for _, group := range res.Groups {
		_, err := s.db.Exec(context.Background(), `DELETE FROM stats
	    WHERE slot_id = $1 AND banner_id = $2 AND group_id = $3`,
			res.SlotID, res.BannerID, group)
		if err != nil {
			return fmt.Errorf("query deletion banner from stats error: %w", err)
		}
	}

	return nil
}

func (s *Storage) AddClick(bannerID, slotID, groupID int) error {
	res := dbResponse{}
	// check if banner and slot and group exist
	err := s.db.QueryRow(
		context.Background(),
		`SELECT banners.id AS "bannerId", slots.id AS "slotId", groups.id AS "groupId"
         FROM banners
         JOIN slots ON slots.id = $2
         JOIN groups ON groups.id = $3
         WHERE banners.id = $1`,
		bannerID, slotID, groupID).Scan(&res.BannerID, &res.SlotID, &res.GroupID)
	if err != nil {
		return fmt.Errorf("query error: %w", err)
	}

	_, err = s.db.Exec(context.Background(), `UPDATE stats
        SET click = click + 1
        WHERE slot_id = $1 AND banner_id = $2 AND group_id = $3`, res.SlotID, res.BannerID, res.GroupID)
	if err != nil {
		return fmt.Errorf("query updating click error: %w", err)
	}

	return nil
}

func (s *Storage) DisplayBanner(slotID, groupID int) (ucb.BannerID, error) {
	res := dbResponse{}
	// check if banner and slot exist
	err := s.db.QueryRow(
		context.Background(),
		`SELECT
           slots.id AS "slotId",
           groups.id AS "groupsId",
           slots.banners AS "banners",
           slots.total_displays AS "totalDisplays"
         FROM slots
         JOIN groups ON groups.id = $2
         WHERE slots.id = $1`,
		slotID, groupID).Scan(&res.SlotID, &res.GroupID, &res.Banners, &res.TotalDisplays)
	if err != nil {
		return 0, fmt.Errorf("query error: %w", err)
	}

	stat := ucb.BannerStat{}
	stats := make(map[ucb.BannerID]ucb.BannerStat)
	for _, banner := range res.Banners {
		err := s.db.QueryRow(context.Background(), `SELECT stats.display, stats.click FROM stats
          WHERE stats.slot_id = $1 AND stats.group_id = $2 AND banner_id = $3`,
			res.SlotID, res.GroupID, banner).Scan(&stat.Trials, &stat.Reward)
		if err != nil {
			return 0, fmt.Errorf("query error: %w", err)
		}
		stats[ucb.BannerID(banner)] = stat
	}

	decision := ucb.MakeDecision(stats, res.TotalDisplays)

	// update total banner display counter for slot
	_, err = s.db.Exec(context.Background(), `UPDATE slots
        SET total_displays = total_displays + 1 WHERE id = $1`, res.SlotID)
	if err != nil {
		if err != nil {
			return 0, fmt.Errorf("query error: %w", err)
		}
	}

	_, err = s.db.Exec(context.Background(), `UPDATE stats
        SET display = display + 1
        WHERE slot_id = $1 AND group_id = $2 AND banner_id = $3`, res.SlotID, res.GroupID, decision)
	if err != nil {
		if err != nil {
			return 0, fmt.Errorf("query error: %w", err)
		}
	}

	return decision, nil
}
