package domain

type Service interface {
	Find(code string) (*Product, error)
	Store(product *Product) error
	Update(product *Product) error
	FindAll() ([]*Product, error)
	Delete(code string) error
}

type Repository interface {
	Find(code string) (*Product, error)
	Store(product *Product) error
	Update(product *Product) error
	FindAll() ([]*Product, error)
	Delete(code string) error
}
