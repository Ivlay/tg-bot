package repository

import (
	"database/sql"
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

func (r *UserSql) Create(user tgbot.User) (int, error) {
	var id int
	query := fmt.Sprintf("insert into %s (userName, firstName, chat_id, user_id) values ($1, $2, $3, $4) returning id", usersTable)
	row := r.db.QueryRow(query, user.UserName, user.FirstName, user.ChatId, user.UserId)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserSql) GetByUserId(id int) (tgbot.User, error) {
	var user tgbot.User

	query := fmt.Sprintf("select * from %s where user_id=$1", usersTable)

	err := r.db.Get(&user, query, id)

	return user, err
}

func (r *UserSql) FindOrCreate(user tgbot.User) (int, error) {
	u, err := r.GetByUserId(user.UserId)
	switch err {
	case sql.ErrNoRows:
		fmt.Printf("User not found, try to create\n %s", err.Error())
		return r.Create(user)
	default:
		return u.UserId, err
	}
}

func (r *UserSql) CreateSubscription() {

}
