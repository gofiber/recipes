package domain

type Book struct {
	Title string `json:"title"`
}

type BooksResponse struct {
	Books []Book `json:"books"`
}
