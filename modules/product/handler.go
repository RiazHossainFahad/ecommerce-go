package product

import (
	"ecommerce/middleware"
	"ecommerce/repo"
)

type Handler struct {
	productRepo      repo.ProductRepo
	midlewareHandler *middleware.MiddlewareHandler
}

func NewHandler(
	productRepo repo.ProductRepo,
	midlewareHandler *middleware.MiddlewareHandler,
) *Handler {
	return &Handler{
		productRepo:      productRepo,
		midlewareHandler: midlewareHandler,
	}
}
