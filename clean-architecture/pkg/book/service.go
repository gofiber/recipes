package book

import (
	"clean-architecture/api/presenter"
	"clean-architecture/pkg/entities"
)

//Service is an interface from which our api module can access our repository of all our models
type Service interface {
	InsertBook(book *entities.Book) (*entities.Book, error)
	FetchBooks() (*[]presenter.Book, error)
	UpdateBook(book *entities.Book) (*entities.Book, error)
	RemoveBook(ID string) error
}

type service struct {
	repository Repository
}

//NewService is used to create a single instance of the service
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

//InsertBook is a service layer that helps insert book in BookShop
func (s *service) InsertBook(book *entities.Book) (*entities.Book, error) {
	return s.repository.CreateBook(book)
}

//FetchBooks is a service layer that helps fetch all books in BookShop
func (s *service) FetchBooks() (*[]presenter.Book, error) {
	return s.repository.ReadBook()
}

//UpdateBook is a service layer that helps update books in BookShop
func (s *service) UpdateBook(book *entities.Book) (*entities.Book, error) {
	return s.repository.UpdateBook(book)
}

//RemoveBook is a service layer that helps remove books from BookShop
func (s *service) RemoveBook(ID string) error {
	return s.repository.DeleteBook(ID)
}
