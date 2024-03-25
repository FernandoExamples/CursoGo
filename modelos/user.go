package modelos

import "time"

type User struct {
	Id        int
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time
	Status    bool
}

func (u *User) SetData(id int, username, password, email string, createdAt time.Time, status bool) {
	u.Id = id
	u.Username = username
	u.Password = password
	u.Email = email
	u.CreatedAt = createdAt
	u.Status = status
}
