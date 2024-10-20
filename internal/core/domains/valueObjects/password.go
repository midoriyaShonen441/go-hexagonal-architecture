package valueObjects

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

type Password struct {
	HashedValue string
}

const (
	MinLength = 8
	MaxLength = 32
)

// List of error type
var (
	ErrTooShort = errors.New(
		fmt.Sprintf("password must have at least %d characters", MinLength))
	ErrTooLong = errors.New(
		fmt.Sprintf("password must not over %d characters", MaxLength))
	ErrNoLowercase = errors.New(
		"password must contain at least one lowercase letter")
	ErrNoUppercase = errors.New(
		"password must contain at least one uppercase letter")
	ErrNoNumber = errors.New(
		"password must contain at least one number")
)

func NewPassword() *Password {
	return &Password{}
}

func (p *Password) IsValid(rawPassword string) error {
	// check password length
	if len(rawPassword) < MinLength {
		return ErrTooShort
	}
	if len(rawPassword) > MaxLength {
		return ErrTooLong
	}

	// check necessary characters
	if !regexp.MustCompile(`[a-z]`).MatchString(rawPassword) {
		return ErrNoLowercase
	}
	if !regexp.MustCompile(`[A-Z]`).MatchString(rawPassword) {
		return ErrNoUppercase
	}
	if !regexp.MustCompile(`[0-9]`).MatchString(rawPassword) {
		return ErrNoNumber
	}

	return nil
}

// Encrypt raw password
func (p *Password) encrypt(rawPassword string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 14)
	if err != nil {
		return "", err
	}

	stringPassword := string(bytes)
	return stringPassword, nil
}

func (p *Password) SetHashedValue(rawPassword string) error {
	hashedPassword, err := p.encrypt(rawPassword)
	if err != nil {
		return err
	}

	p.HashedValue = hashedPassword

	return nil
}

// Verify password correctness
func (p *Password) Verify(rawPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.HashedValue), []byte(rawPassword))
	return err == nil
}
