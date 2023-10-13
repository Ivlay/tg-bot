package service

import (
	"log"
	"math/rand"

	tgbot "github.com/Ivlay/go-telegram-bot"
	"github.com/Ivlay/go-telegram-bot/pkg/htmlParser"
	"github.com/Ivlay/go-telegram-bot/pkg/repository"
)

type ProductService struct {
	repo   repository.Product
	parser *htmlParser.HtmlParser
}

func NewProductService(repo repository.Product, parser *htmlParser.HtmlParser) *ProductService {
	return &ProductService{repo: repo, parser: parser}
}

func (p *ProductService) Create() {
	p.repo.Create()
}

func (p *ProductService) GetAllProducts() ([]tgbot.Product, error) {
	return p.repo.GetAllProducts()
}

func (p *ProductService) GetByIdsWithUser(productIds []int) ([]tgbot.UserWithProducts, error) {
	return p.repo.GetByIdsWithUser(productIds)
}

func (p *ProductService) Update() ([]int, error) {
	products := []tgbot.Product{
		{Title: `Pro 16" M1 Pro 10-Core/ 16GB/ 512GB SSD 16-core GPU`, PriceId: "macbook_pro_16_2021_", Price: rand.Intn(120000)},
		{Title: `Pro 16" M1 Max 10-Core/ 32GB/ 1TB SSD 32-core GPU`, PriceId: "macbook_pro_16_2021_", Price: rand.Intn(170000)},
	}

	return p.repo.Update(products)
}

func (p *ProductService) Prepare() {
	count, err := p.repo.GetCount()
	if err != nil {
		log.Fatalf("error while get count of products %s", err.Error())
		return
	}

	if count <= 0 {
		pp := p.parser.PrepareProducts()
		err := p.repo.Prepare(pp)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (r *ProductService) GetByUserId(userId int) ([]tgbot.Product, error) {
	return r.repo.GetByUserId(userId)
}

func (r *ProductService) GetProductsListByProductId() ([]tgbot.ProductSubscriptions, error) {
	return r.repo.GetProductsListSubscriptions()
}
