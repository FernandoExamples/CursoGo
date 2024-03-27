package models

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

var Users = []User{
	{Id: "1", Name: "John Doe", IsActive: true},
	{Id: "2", Name: "Jane Doe", IsActive: false},
	{Id: "3", Name: "John Smith", IsActive: true},
	{Id: "4", Name: "Jane Smith", IsActive: false},
	{Id: "5", Name: "John Brown", IsActive: true},
}
