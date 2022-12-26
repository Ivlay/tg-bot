package repository

import (
	"fmt"

	tgbot "github.com/Ivlay/go-telegram-bot"
	"github.com/jmoiron/sqlx"
)

type ProductSql struct {
	db *sqlx.DB
}

func NewProductSql(db *sqlx.DB) *ProductSql {
	return &ProductSql{
		db: db,
	}
}

func (r *ProductSql) Create() {

}

func (r *ProductSql) Update() {

}

func (r *ProductSql) GetByUserId() {

}

func (r *ProductSql) Prepare(products []tgbot.Product) error {
	query := fmt.Sprintf(`
		insert into %s (price_id, title, price)
		values (:price_id, :title, :price)`,
		productsTable,
	)
	res, err := r.db.NamedExec(query, products)
	if err != nil {
		return err
	}

	fmt.Println("Res", res)
	return nil
}
