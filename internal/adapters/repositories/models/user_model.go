package models

import "github.com/midoriyaShonen441/hexagonal-architecture/internal/core/domains/entities"

type UserModel struct {
	Id       uint `gorm:"primary_key"`
	Name     string
	Email    string `gorm:"unique"`
	Username string `gorm:"unique"`
	Password string
}

func NewUserModel(u *entities.User) *UserModel {
	return &UserModel{
		Id:       0,
		Name:     u.Name,
		Email:    u.Email.Value,
		Username: u.Username,
		Password: u.Password.HashedValue,
	}
}
