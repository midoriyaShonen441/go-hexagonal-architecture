package services

import (
	"github.com/midoriyaShonen441/hexagonal-architecture/internal/core/domains/entities"
	"github.com/midoriyaShonen441/hexagonal-architecture/internal/core/ports"
)

type AuthService struct {
	authRepository ports.AuthRepository
}

var _ ports.AuthService = (*AuthService)(nil)

func NewAuthService(authRepository ports.AuthRepository) *AuthService {
	return &AuthService{
		authRepository: authRepository,
	}
}

func (a AuthService) Login(username string, password string) (*entities.User, error) {
	user, err := a.authRepository.FindByUsername(username)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a AuthService) Register(user *entities.User) error {
	return a.authRepository.Create(user)
}
