package repository

import (
	tgbot "github.com/Ivlay/go-telegram-bot"
	"github.com/jmoiron/sqlx"
)

type User interface {
	CreateUser(user tgbot.User) (int, error)
	GetUserByUserId(id int) (tgbot.User, error)
	FindOrCreateUser(user tgbot.User) (int, error)
}

type Repository struct {
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserSql(db),
	}
}