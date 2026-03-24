package auth

import "ecommerce/domain"

type Service interface {
	Find(email, password string) (*domain.User, error)
}
