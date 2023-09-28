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

func (r *ProductSql) GetByUserId(userId int) ([]tgbot.Product, error) {
	var pp []tgbot.Product

	query := fmt.Sprintf(`
		select p.id as id, p.title as title, p.price as price, p.price_id as price_id, p.updated_at as updated_at from %s pl inner join %s p on pl.product_id = p.id
		where pl.user_id = %d
		order by pl.created_at
	`, productsLists, productsTable, userId)

	err := r.db.Select(&pp, query)

	return pp, err
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
