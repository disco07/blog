package repository

import "github.com/jmoiron/sqlx"

type Repository struct {
	DB *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return Repository{
		DB: db,
	}
}
