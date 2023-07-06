package fixtures

import (
	"context"

	"app/config/database"
	"app/entity"
)

func LoadTodos(name string) (err error) {
	client, err := database.Connect(name)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	bulk := make([]*entity.TodoCreate, len(todos))
	for i, todo := range todos {
		bulk[i] = client.Todo.Create().SetContent(todo.Content).SetCompleted(todo.Completed)
	}
	_, err = client.Todo.CreateBulk(bulk...).Save(context.Background())
	return
}

var todos = []entity.Todo{
	{
		Content:   "Buy tickets for a rock concert",
		Completed: true,
	},
	{
		Content:   "Explore a hidden beach",
		Completed: false,
	},
	{
		Content:   "Learn to play the guitar",
		Completed: false,
	},
	{
		Content:   "Try a new recipe and cook a gourmet meal",
		Completed: false,
	},
	{
		Content:   "Go skydiving and feel the adrenaline rush",
		Completed: false,
	},
	{
		Content:   "Visit famous landmarks in a foreign country",
		Completed: false,
	},
	{
		Content:   "Write a novel and become a bestselling author",
		Completed: false,
	},
	{
		Content:   "Learn to speak a new language fluently",
		Completed: false,
	},
	{
		Content:   "Start a successful online business",
		Completed: false,
	},
	{
		Content:   "Run a marathon and achieve a personal record",
		Completed: false,
	},
}
