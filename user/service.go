package user

import "ecommerce/domain"

type service struct {
	usrRepo UserRepo
}

func NewService(usrRepo UserRepo) Service {
	return &service{
		usrRepo: usrRepo,
	}
}

func (svc service) EmptyUser() domain.User {
	return svc.usrRepo.EmptyUser()
}

func (svc service) List() ([]*domain.User, error) {
	return svc.usrRepo.List()
}

func (svc service) Store(usr domain.User) (*domain.User, error) {
	return svc.usrRepo.Store(usr)
}

func (svc service) Find(email, password string) (*domain.User, error) {
	return svc.usrRepo.Find(email, password)
}
