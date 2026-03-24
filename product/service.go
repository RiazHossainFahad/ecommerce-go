package product

import "ecommerce/domain"

type service struct {
	productRepo ProductRepo
}

func NewService(productRepo ProductRepo) Service {
	return &service{
		productRepo: productRepo,
	}
}

func (svc service) EmptyProduct() domain.Product {
	return svc.productRepo.EmptyProduct()
}

func (svc service) List() ([]*domain.Product, error) {
	productList, err := svc.productRepo.List()
	if err != nil {
		return nil, err
	}

	return productList, nil
}

func (svc service) Store(reqProduct domain.Product) (*domain.Product, error) {
	product, err := svc.productRepo.Store(reqProduct)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (svc service) Get(id int) (*domain.Product, error) {
	product, err := svc.productRepo.Get(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (svc service) Update(id int, reqProduct domain.Product) (*domain.Product, error) {
	product, err := svc.productRepo.Update(id, reqProduct)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (svc service) Delete(id int) error {
	err := svc.productRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (svc service) Exists(id int) (bool, error) {
	isExist, err := svc.productRepo.Exists(id)
	if err != nil {
		return isExist, err
	}

	return isExist, nil
}
