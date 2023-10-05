package tgbot

type User struct {
	Id        int    `json:"id" db:"id"`
	ChatId    int64  `json:"chatId" db:"chat_id"`
	UserId    int    `json:"userId" db:"user_id"`
	FirstName string `json:"firstName" db:"firstname"`
	UserName  string `json:"userName" db:"username"`
	CreatedAt string `json:"-" db:"created_at"`
}

type UserList struct {
	Id        int
	UserId    int
	ProductId string
}

type UserWithProducts struct {
	UserId   int    `json:"userId" db:"user_id"`
	UserName string `json:"userName" db:"username"`
	ChatId   int64  `json:"chatId" db:"chat_id"`
	Products []Product
}
