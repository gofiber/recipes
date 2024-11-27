---
title: Entgo ORM (MySQL)
keywords: [ent, mysql, orm, rest]
description: Using Entgo ORM with MySQL
---

# Example ent ORM for fiber with MySQL

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/ent-mysql) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/ent-mysql)

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
