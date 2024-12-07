package database

import "context"

// This is just an example and not for production use
func newMemoryDB() Database {
	return &memoryDB{
		records:   make([]Book, 0, 10),
		idCounter: 0,
	}
}

type memoryDB struct {
	records   []Book
	idCounter int
}

func (db *memoryDB) LoadAllBooks(_ context.Context) ([]Book, error) {
	return db.records, nil
}

func (db *memoryDB) CreateBook(_ context.Context, newBook NewBook) error {
	db.records = append(db.records, Book{
		ID:    db.idCounter,
		Title: newBook.Title,
	})
	db.idCounter++
	return nil
}

func (db *memoryDB) CloseConnections() {
}
