package service

import (
	tgbot "github.com/Ivlay/go-telegram-bot"
	"github.com/Ivlay/go-telegram-bot/pkg/repository"
)

type User interface {
	CreateUser(user tgbot.User) (int, error)
	GetUserByUserId(id int) (tgbot.User, error)
}

type Service struct {
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repos.User),
	}
}