# Clean Architecture with GoFiber

A simple book shop CRUD demonstration implementing the Clean Architecture in GoFiber

## Installation

In order to run the project, please follow the following steps:

1. Clone the Repo
2. Go to the "clean-architecture" folder
3. Run `go get`
4. Replace your Mongo DB Connection string in `app.go`

## Routes

|  API Path  | Method |               What it does              |
|:----------:|:------:|:---------------------------------------:|
| /api/books |   GET  | Fetches the list of books from the shop |
| /api/books |  POST  |      Creates/Adds book to the shop      |
| /api/books | DELETE |    Removes/Deletes book from the shop   |
| /api/books |  PUT |  Updates the book details from the shop |
