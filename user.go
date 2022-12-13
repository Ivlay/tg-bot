package tgbot

type User struct {
	Id        int    `json:"-" db:"id"`
	ChatId    int64  `json:"chatId" db:"chat_id"`
	UserId    int    `json:"userId" db:"user_id"`
	FirstName string `json:"firstName" db:"firstname"`
	UserName  string `json:"userName" db:"username"`
}