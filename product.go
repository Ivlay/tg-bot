package tgbot

type Product struct {
	Id        int    `json:"id" db:"id"`
	Title     string `json:"title" db:"title"`
	PriceId   string `json:"priceId" db:"price_id"`
	Price     int    `json:"price" db:"price"`
}