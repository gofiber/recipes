package types

// TodoResponse struct contains the todo field which should be returned in a response
type TodoResponse struct {
	ID        uint   `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

// CreateDTO struct defines the /todo/create payload
type CreateDTO struct {
	Task string `json:"task" validate:"required,min=3,max=150"`
}

// TodoCreateResponse struct defines the /todo/create response
type TodoCreateResponse struct {
	Todo *TodoResponse `json:"todo"`
}

// TodosResponse defines the todos list
type TodosResponse struct {
	Todos *[]TodoResponse `json:"todos"`
}

// CheckTodoDTO defined the payload for the check todo
type CheckTodoDTO struct {
	Completed bool `json:"completed"`
}
