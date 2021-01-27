package storage

import sqlstorage "github.com/ivfisunov/banner-rotator/internal/storage/sql"

type Slot struct {
	ID          int
	Description string
	Banners     []int
}

type Banner struct {
	ID          int
	Description string
}

type Group struct {
	ID          int
	Description string
}

type Storage interface {
	Connect() error
	Close() error
	AddBanner(idBanner, idSlot int) error
	DeleteBanner(idBanner, idSlot int) error
	AddClick(idBunner, idSlot, idGroup int) error
}

func New(dsn string) (Storage, error) {
	stor, err := sqlstorage.New(dsn)
	if err != nil {
		return nil, err
	}

	return stor, nil
}
