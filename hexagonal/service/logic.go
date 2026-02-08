package service

import "catalog/domain"

type service struct {
	productRepo domain.Repository
}

// NewProductService returns a domain.Service backed by the provided domain.Repository.
func NewProductService(productRepo domain.Repository) domain.Service {
	return &service{productRepo: productRepo}
}

func (s *service) Find(code string) (*domain.Product, error) {
	return s.productRepo.Find(code)
}

func (s *service) Store(product *domain.Product) error {
	return s.productRepo.Store(product)
}

func (s *service) Update(product *domain.Product) error {
	return s.productRepo.Update(product)
}

func (s *service) FindAll() ([]*domain.Product, error) {
	return s.productRepo.FindAll()
}

func (s *service) Delete(code string) error {
	return s.productRepo.Delete(code)
}
