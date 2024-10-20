package valueObjects

import (
	"errors"
	"regexp"
)

type Email struct {
	Value string
}

var InvalidEmailError = errors.New("invalid email format")

func isValidEmail(email string) error {
	if !regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(email) {
		return InvalidEmailError
	}

	return nil
}

func NewEmail(email string) (*Email, error) {
	if err := isValidEmail(email); err != nil {
		return nil, err
	}

	return &Email{Value: email}, nil
}
