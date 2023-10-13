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
	GetAllProducts() ([]tgbot.Product, error)
	GetByIdsWithUser(productIds []int) ([]tgbot.UserWithProducts, error)
	GetByUserId(userId int) ([]tgbot.Product, error)
	GetProductsListSubscriptions() ([]tgbot.ProductSubscriptions, error)
	Prepare(products []tgbot.Product) error
	Update(products []tgbot.Product) ([]int, error)
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
