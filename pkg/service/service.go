package service

import (
	tgbot "github.com/Ivlay/go-telegram-bot"
	"github.com/Ivlay/go-telegram-bot/pkg/htmlParser"
	"github.com/Ivlay/go-telegram-bot/pkg/repository"
)

type User interface {
	Create(user tgbot.User) (int, error)
	GetByUserId(id int) (tgbot.User, error)
	FindOrCreate(user tgbot.User) (int, error)
}

type Product interface {
	Create()
	GetByUserId()
	Update()
	Prepare()
}

type Parser interface {
	FindNodes()
}

type Service struct {
	User
	Product
}

func NewService(repos *repository.Repository, parser *htmlParser.HtmlParser) *Service {
	return &Service{
		User:    NewUserService(repos.User),
		Product: NewProductService(repos.Product, parser),
	}
}
