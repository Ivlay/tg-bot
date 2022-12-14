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

func (r *UserService) Create(user tgbot.User) (int, error) {
	return r.repo.Create(user)
}

func (r *UserService) GetByUserId(id int) (tgbot.User, error) {
	return r.repo.GetByUserId(id)
}

func (r *UserService) FindOrCreate(user tgbot.User) (int, error) {
	return r.repo.FindOrCreate(user)
}

func (r *UserService) CreateSubscription() {

}
