package product

import "ecommerce/domain"

type Service interface {
	EmptyProduct() domain.Product
	List(page, limit int) ([]*domain.Product, error)
	Store(domain.Product) (*domain.Product, error)
	Get(int) (*domain.Product, error)
	Update(int, domain.Product) (*domain.Product, error)
	Delete(int) error
	Unique(string, int) (bool, error)
	Count() (int, error)
}
