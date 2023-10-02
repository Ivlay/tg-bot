package repository

import (
	tgbot "github.com/Ivlay/go-telegram-bot"
	"github.com/jmoiron/sqlx"
)

type User interface {
	Create(user tgbot.User) (int, error)
	GetByUserId(id int) (tgbot.User, error)
	FindOrCreate(user tgbot.User) (int, error)
}

type Product interface {
	Create()
	Update()
	GetByUserIds(ids []int) ([]tgbot.Product, error)
	Prepare(products []tgbot.Product) error
	UpdateProducts(products []tgbot.Product) ([]int, error)
	GetCount() (int64, error)
}

type Repository struct {
	User
	Product
}

func New(db *sqlx.DB) *Repository {
	return &Repository{
		User:    NewUserSql(db),
		Product: NewProductSql(db),
	}
}
