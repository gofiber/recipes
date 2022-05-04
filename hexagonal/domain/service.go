package domain

/* A service interface defining all CRUD operations*/
type Service interface {
	Find(code string) (*Product, error)
	Store(product *Product) error
	Update(product *Product) error
	FindAll() ([]*Product, error)
	Delete(code string) error
}

/* a repository interface acting like a port for the database implementation*/
type Repository interface {
	Find(code string) (*Product, error)
	Store(product *Product) error
	Update(product *Product) error
	FindAll() ([]*Product, error)
	Delete(code string) error
}
