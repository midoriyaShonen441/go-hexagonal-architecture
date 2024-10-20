package ports

import (
	"github.com/gofiber/fiber/v2"
	"github.com/midoriyaShonen441/hexagonal-architecture/internal/core/domains/entities"
)

type (
	AuthService interface {
		Login(username string, password string) (*entities.User, error)
		Register(user *entities.User) error
	}

	AuthRepository interface {
		FindByUsername(username string) (*entities.User, error)
		Create(user *entities.User) error
	}

	AuthHandler interface {
		Login(c *fiber.Ctx) error
		Register(c *fiber.Ctx) error
	}
)
