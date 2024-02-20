package repository

import "github.com/jmoiron/sqlx"

func (a *Repository) CreateUserWithTgId(tgId string) error {
	// logic will here
	return nil
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

type Repository struct {
	db *sqlx.DB
}
