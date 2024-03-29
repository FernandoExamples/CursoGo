package models

type Todo struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var Todos = []Todo{
	{ID: 1, Name: "Task 1", Completed: false},
	{ID: 2, Name: "Task 2", Completed: false},
	{ID: 3, Name: "Task 3", Completed: false},
}
