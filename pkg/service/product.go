package service

import (
	"github.com/Ivlay/go-telegram-bot/pkg/repository"
)

type ProductService struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) *ProductService {
	return &ProductService{repo: repo}
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
	r.repo.Prepare()
}