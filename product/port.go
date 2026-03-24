package product

import (
	"ecommerce/domain"
	prdctHandler "ecommerce/modules/product"
)

type Service interface {
	prdctHandler.Service
}

type ProductRepo interface {
	EmptyProduct() domain.Product

	List(page, limit int) ([]*domain.Product, error)
	Store(domain.Product) (*domain.Product, error)
	Get(int) (*domain.Product, error)
	Update(int, domain.Product) (*domain.Product, error)
	Delete(int) error
	Exists(int) (bool, error)
	Unique(string, int) (bool, error)
	Count() (int, error)
}
