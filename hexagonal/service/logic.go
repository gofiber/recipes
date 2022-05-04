package service

import "catalog/domain"

type service struct {
	productRepo domain.Repository
}

// NewProductService a service struct that attaches to a repository via the Repository interface
func NewProductService(productRepo domain.Repository) *service {
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
