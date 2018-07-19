package handmedown

import "time"

type UserId int

type User struct {
	ID        UserId
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserService interface {
	CreateUser(user *User) error
	FindUserByID(id UserId) (*User, error)
}
