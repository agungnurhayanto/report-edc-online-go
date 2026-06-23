package monitoring

import (
	"context"
	"database/sql"
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
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}
