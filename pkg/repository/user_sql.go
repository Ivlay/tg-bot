package repository

import (
	"fmt"

	tgbot "github.com/Ivlay/go-telegram-bot"
	"github.com/jmoiron/sqlx"
)

type UserSql struct {
	db *sqlx.DB
}

func NewUserSql(db *sqlx.DB) *UserSql {
	return &UserSql{db: db}
}

func (r* UserSql) CreateUser(user tgbot.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (userName, firstName, chat_id, user_id) values ($1, $2, $3, $4) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.UserName, user.FirstName, user.ChatId, user.UserId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil;
}

func (r* UserSql) GetUserByUserId(id int) (tgbot.User, error) {
	var user tgbot.User

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id=$1", usersTable)

	err := r.db.Get(&user, query, id)

	return user, err
}
