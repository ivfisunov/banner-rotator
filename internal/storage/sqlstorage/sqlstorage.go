package sqlstorage

import (
	"github.com/ivfisunov/banner-rotator/internal/storage/stortypes"
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	dsn string
	db  *sqlx.DB
}

func New(dsn string) (*Storage, error) {
	return &Storage{dsn: dsn}, nil
}

func (s *Storage) Connect() error {
	// setup postgres connection
	db, err := sqlx.Connect("postgres", s.dsn)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db

	return nil
}

func (s *Storage) Close() error {
	err := s.db.Close()
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) AddBanner(bannerID, slotID int) error {
	return nil
}

func (s *Storage) DeleteBanner(bannerID, slotID int) error {
	return nil
}

func (s *Storage) AddClick(bannerID, slotID, groupID int) error {
	return nil
}

func (s *Storage) DisplayBanner(slotID, groupID int) (stortypes.Banner, error) {
	return stortypes.Banner{}, nil
}
