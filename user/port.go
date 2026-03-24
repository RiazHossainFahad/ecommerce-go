package user

import (
	"ecommerce/domain"
	usrHandler "ecommerce/modules/user"
)

type Service interface {
	usrHandler.Service
}

type UserRepo interface {
	EmptyUser() domain.User

	List() ([]*domain.User, error)
	Store(domain.User) (*domain.User, error)
	Find(email, password string) (*domain.User, error)
}
