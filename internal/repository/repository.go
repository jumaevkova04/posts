package repository

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository struct {
	*sqlx.DB
}

func NewRepository(dsn string) (*Repository, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return &Repository{db}, nil
}
