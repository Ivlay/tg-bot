package repository

import (
	"fmt"
	"strings"

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

func (r *ProductSql) GetAllProducts() ([]tgbot.Product, error) {
	var pp []tgbot.Product

	query := fmt.Sprintf(`
		select * from %s
	`, productsTable)

	err := r.db.Select(&pp, query)

	return pp, err
}

func (r *ProductSql) Update(products []tgbot.Product) ([]int, error) {
	var ids []int

	query := fmt.Sprintf(`
		update %s
			set price = :price
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

func (r *ProductSql) GetByIdsWithUser(productIds []int) ([]tgbot.UserWithProducts, error) {
	ids := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(productIds)), ", "), "[]")

	query1 := fmt.Sprintf(`
		select u.chat_id, u.user_id, u.username, p.title, p.price, p.old_price, p.updated_at
		from %s pl
		join %s p on pl.product_id = p.id
		join %s u on pl.user_id = u.user_id
		where pl.product_id = any(array[%s])
	`, productsListsTable, productsTable, usersTable, ids)

	rows, err := r.db.Query(query1)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	usersMap := make(map[int]*tgbot.UserWithProducts)

	for rows.Next() {
		var user tgbot.UserWithProducts
		var product tgbot.Product

		err := rows.Scan(&user.ChatId, &user.UserId, &user.UserName, &product.Title, &product.Price, &product.OldPrice, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}

		if existingUser, ok := usersMap[user.UserId]; ok {
			existingUser.Products = append(existingUser.Products, product)
		} else {
			user.Products = []tgbot.Product{product}
			usersMap[user.UserId] = &user
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	usersWithProduct := make([]tgbot.UserWithProducts, 0, len(usersMap))
	for _, user := range usersMap {
		usersWithProduct = append(usersWithProduct, *user)
	}

	return usersWithProduct, nil
}

func (r *ProductSql) GetByUserId(userId int) ([]tgbot.Product, error) {
	var pp []tgbot.Product

	query := fmt.Sprintf(`
		select p.title, p.price, p.updated_at, p.price
		from %s pl
		join %s p on pl.product_id = p.id
		where pl.user_id = %d
		order by pl.created_at
	`, productsListsTable, productsTable, userId)

	err := r.db.Select(&pp, query)

	return pp, err
}

func (r *ProductSql) UpdateProducts(products []tgbot.Product) ([]int, error) {
	var ids []int

	query := fmt.Sprintf(`
		update %s
			set price = :price
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

func (r *ProductSql) GetProductsListSubscriptions() ([]tgbot.ProductSubscriptions, error) {
	var ps []tgbot.ProductSubscriptions

	query := fmt.Sprintf(`
		select p.title, count(p) as subscribers
		from %s pl
		join %s p on pl.product_id = p.id
		group by p.title
	`, productsListsTable, productsTable)

	err := r.db.Select(&ps, query)

	return ps, err
}
