package repository

import (
	"github.com/jmoiron/sqlx"
	"log"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (repo *Repository) CreateUserWithTgId(tgId int64) error {

	request := `INSERT INTO users(telegram_id) VALUES($1);`

	if _, err := repo.db.Exec(request, tgId); err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
