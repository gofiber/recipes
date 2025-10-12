package services

import "auth-jwt-gorm/models"

// AuthService provides authentication functionality
type ProductService struct {
	productRepo *models.ProductRepository
}

// NewAuthService creates a new authentication service
func NewProductService(productRepo *models.ProductRepository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}
