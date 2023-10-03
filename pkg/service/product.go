package service

import (
	"log"

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

func (p *ProductService) GetByUserIds(id []int) ([]tgbot.Product, error) {
	return p.repo.GetByUserIds(id)
}

func (p *ProductService) Update() {
	p.repo.Update()
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

func (p *ProductService) UpdateProducts() ([]int, error) {
	// products := p.parser.PrepareProducts()

	products := []tgbot.Product{
		{Title: `Pro 14" M1 Pro 8-Core/ 16GB/ 512GB SSD 14-core GPU`, PriceId: "macbook_pro_14_2021_", Price: 50000},
		{Title: `Pro 16" M1 Max 10-Core/ 32GB/ 1TB SSD 32-core GPU`, PriceId: "macbook_pro_16_2021_", Price: 60000},
	}

	return p.repo.UpdateProducts(products)
}

func (p *ProductService) GetProductsListByProductId([]int) {

}
