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

func (r *ProductService) Create() {
	r.repo.Create()
}

func (r *ProductService) GetByUserId(id int) ([]tgbot.Product, error) {
	return r.repo.GetByUserId(id)
}

func (r *ProductService) Update() {
	r.repo.Update()
}

func (r *ProductService) Prepare() {
	pp := r.parser.PrepareProducts()
	err := r.repo.Prepare(pp)
	if err != nil {
		log.Fatal(err)
	}
}
