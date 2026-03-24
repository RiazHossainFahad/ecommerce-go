package user

import "ecommerce/domain"

type Service interface {
	EmptyUser() domain.User

	List() ([]*domain.User, error)
	Store(domain.User) (*domain.User, error)
	Find(email, password string) (*domain.User, error)
}
