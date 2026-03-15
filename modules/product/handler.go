package product

import "ecommerce/repo"

type Handler struct {
	productRepo repo.ProductRepo
}

func NewHandler(productRepo repo.ProductRepo) *Handler {
	return &Handler{
		productRepo: productRepo,
	}
}
