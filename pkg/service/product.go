package service

import (
	"log"

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

func (r *ProductService) GetByUserId() {
	r.repo.GetByUserId()
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
