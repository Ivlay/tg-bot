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

func (r *ProductSql) GetByUserIds(userIds []int) ([]tgbot.Product, error) {
	var pp []tgbot.Product

	query := fmt.Sprintf(`
		select p.id as id, p.title as title, p.price as price, p.price_id as price_id, p.updated_at as updated_at from %s pl inner join %s p on pl.product_id = p.id
		where pl.user_id in (%d)
		order by pl.created_at
	`, productsLists, productsTable, userIds)

	err := r.db.Select(&pp, query)

	return pp, err
}

func (r *ProductSql) UpdateProducts(products []tgbot.Product) ([]int, error) {
	var ids []int

	query := fmt.Sprintf(`
		update %s
			set price = :price, old_price = price
			where price_id = :price_id and title = :title and price != :price
			returning id
	`, productsTable)

	for _, product := range products {
		rows, err := r.db.NamedQuery(query, product)
		if err != nil {
			return ids, err
		}
		defer rows.Close()
		for rows.Next() {
			var id int
			err := rows.Scan(&id)
			if err != nil {
				return ids, err
			}

			ids = append(ids, id)
		}
	}

	return ids, nil
}

func (r *ProductSql) Prepare(products []tgbot.Product) error {
	query := fmt.Sprintf(`
		insert into %s (price_id, title, price, old_price)
		values (:price_id, :title, :price, :price)`,
		productsTable,
	)
	res, err := r.db.NamedExec(query, products)
	if err != nil {
		return err
	}

	fmt.Println("Res\n", res)
	return nil
}

func (r *ProductSql) GetCount() (int64, error) {
	var rowCount int64
	query := fmt.Sprintf(`
		select count(*) from %s
	`, productsTable)

	err := r.db.Get(&rowCount, query)
	if err != nil {
		return rowCount, err
	}

	return rowCount, nil
}
