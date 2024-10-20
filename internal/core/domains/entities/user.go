package entities

import (
	"github.com/midoriyaShonen441/hexagonal-architecture/internal/core/domains/valueObjects"
)

type User struct {
	Id       uint
	Name     string
	Email    valueObjects.Email
	Username string
	Password valueObjects.Password
}

func NewUser(id uint, name string, email string, username string, password string) (*User, error) {
	voEmail, err := valueObjects.NewEmail(email)
	if err != nil {
		return nil, err
	}
	voPassword := valueObjects.NewPassword()

	// Validate the password
	if err := voPassword.IsValid(password); err != nil {
		return nil, err
	}

	// Encrypt the password
	if err := voPassword.SetHashedValue(password); err != nil {
		return nil, err
	}

	return &User{
		Id:       id,
		Name:     name,
		Email:    *voEmail,
		Username: username,
		Password: *voPassword,
	}, nil
}
