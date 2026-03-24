package user

import (
	"ecommerce/user"
)

type Service struct {
	usrSvc user.Service
}

func NewService(usrSvc user.Service) *Service {
	return &Service{
		usrSvc: usrSvc,
	}
}
