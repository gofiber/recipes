package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Immutable().
			Comment("The unique identifier for the todo."),
		field.
			Text("content").
			Comment("The content of the todo."),
		field.
			Bool("completed").
			Comment("Indicates whether the todo is completed or not."),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return nil
}
