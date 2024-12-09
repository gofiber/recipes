package domain

// Book represents a book
type Book struct {
	Title string `json:"title"`
}

// BooksResponse represents a response containing a list of books
type BooksResponse struct {
	Books []Book `json:"books"`
}
