package repository

import "github.com/jmoiron/sqlx"

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