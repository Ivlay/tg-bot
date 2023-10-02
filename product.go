package tgbot

type Product struct {
	Id        int    `json:"id" db:"id"`
	Title     string `json:"title" db:"title"`
	PriceId   string `json:"priceId" db:"price_id"`
	Price     int    `json:"price" db:"price"`
	UpdatedAt string `json:"updatedAt" db:"updated_at"`
	OldPrice  int    `json:"oldPrice" db:"old_price"`
}

type ProductList struct {
	Id        int    `json:"id" db:"id"`
	ProductId int    `json:"productId" db:"product_id"`
	UserId    int    `json:"userId" db:"user_id"`
	CreatedAt string `json:"createdAt" db:"created_at"`
}
