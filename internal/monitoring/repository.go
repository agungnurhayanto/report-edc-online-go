package monitoring

import (
	"context"
	"database/sql"
	"time"
)

type Repository interface {
	Count() (int, error)

	BulkInsert(
		ctx context.Context,
		data []Monitoring,
	) error

	//FindAll() ([]Monitoring, error)
	FindAll(
		limit int,
		offset int,
	) ([]Monitoring, error)

	// validated
	ExistByDate(tgl time.Time) (bool, error)
	DeleteByDate(ctx context.Context, tgl time.Time) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
