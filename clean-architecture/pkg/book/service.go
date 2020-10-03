package book

import "clean-architecture/pkg/entities"
//Service is an interface from which our api module can access our repository of all our models
type Service interface {
	InsertBook(book *entities.Book) (*entities.Book, error)
	FetchBooks() (*[]entities.Book, error)
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

func (s *service) InsertBook(book *entities.Book) (*entities.Book, error) {
	return s.repository.CreateBook(book)
}
func (s *service) FetchBooks() (*[]entities.Book, error) {
	return s.repository.ReadBook()

}
func (s *service) UpdateBook(book *entities.Book) (*entities.Book, error) {
	return s.repository.UpdateBook(book)
}
func (s *service) RemoveBook(ID string) error {
	return s.repository.DeleteBook(ID)
}
