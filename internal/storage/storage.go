package storage

import (
	"github.com/ivfisunov/banner-rotator/internal/storage/sqlstorage"
	"github.com/ivfisunov/banner-rotator/internal/storage/stortypes"
)

func New(dsn string) (stortypes.Storage, error) {
	stor, err := sqlstorage.New(dsn)
	if err != nil {
		return nil, err
	}

	return stor, nil
}
