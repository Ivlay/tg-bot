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
	GetAllProducts() ([]tgbot.Product, error)
	GetByIdsWithUser(productIds []int) ([]tgbot.UserWithProducts, error)
	GetByUserId(userId int) ([]tgbot.Product, error)
	GetProductsListByProductId() ([]tgbot.ProductSubscriptions, error)
	Update() ([]int, error)
	Prepare()
}

type Parser interface {
	FindNodes()
}

type Service struct {
	User
	Product
}

func New(repos *repository.Repository, parser *htmlParser.HtmlParser) *Service {
	UserService := NewUserService(repos.User)
	ProductService := NewProductService(repos.Product, parser)

	return &Service{
		User:    UserService,
		Product: ProductService,
	}
}
