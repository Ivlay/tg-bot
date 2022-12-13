package service

import (
	tgbot "github.com/Ivlay/go-telegram-bot"
	"github.com/Ivlay/go-telegram-bot/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (r *UserService) CreateUser(user tgbot.User) (int, error) {
	return r.repo.CreateUser(user)
}

func (r *UserService) GetUserByUserId(id int) (tgbot.User, error) {
	return r.repo.GetUserByUserId(id)
}