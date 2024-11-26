---
title: Entgo ORM (MySQL)
keywords: [ent, mysql, orm, rest]
---

# Example ent ORM for fiber with MySQL

A sample program how to connect ent ORM

## How to start (If no ent dir)
Execute command first
```bash
go run -mod=mod entgo.io/ent/cmd/ent new Book
```
go to `./ent/schema/book.go` and add fields(you want) to Book Schema
```go
// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.String("author").NotEmpty(),
	}
}
```
Execute command
```bash
go generate ./ent
```

### Endpoints

| Method | URL         | Description     |
|--------|-------------|-----------------|
| GET    | /book       | All Books Info  |
| GET    | /book:id    | One Book Info   |
| POST   | /create     | One Book Add    |
| PUT    | /update/:id | One Book Update |
| DELETE | /delete/:id | One Book Delete |
