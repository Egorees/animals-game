package repository

import (
	"animals-game/internal/users"
	"github.com/jmoiron/sqlx"
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
		return err
	}
	return nil
}

func (repo *Repository) GetUserIdByTgId(tgId int64) (int, error) {
	var user []users.User

	request := `SELECT id FROM users WHERE (telegram_id = $1)`

	err := repo.db.Select(&user, request, tgId)
	if err != nil {
		return 0, err
	}

	return user[0].Id, nil
}
