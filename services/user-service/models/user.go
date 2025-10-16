package models

type User struct {
	ID gocql
}

func NewUser(id, username, email, password string) *User {
	return &User{
		ID:       id,
		Username: username,
		Email:    email,
		Password: password,
	}
}
