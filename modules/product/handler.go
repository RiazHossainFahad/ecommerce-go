package product

import (
	"ecommerce/middleware"
)

type Handler struct {
	productService   Service
	midlewareHandler *middleware.MiddlewareHandler
}

func NewHandler(
	productService Service,
	midlewareHandler *middleware.MiddlewareHandler,
) *Handler {
	return &Handler{
		productService:   productService,
		midlewareHandler: midlewareHandler,
	}
}
