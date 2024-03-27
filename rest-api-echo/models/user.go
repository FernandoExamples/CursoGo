package models

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique;not null"`
	Tasks     []Task `json:"tasks"`
}
