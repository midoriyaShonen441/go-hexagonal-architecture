package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/midoriyaShonen441/hexagonal-architecture/internal/adapters/handlers/schemas"
	"github.com/midoriyaShonen441/hexagonal-architecture/internal/core/domains/entities"
	"github.com/midoriyaShonen441/hexagonal-architecture/internal/core/ports"
)

type FiberAuthHandler struct {
	authService ports.AuthService
}

func NewFiberAuthHandler(authService ports.AuthService) *FiberAuthHandler {
	return &FiberAuthHandler{
		authService: authService,
	}
}

func (h *FiberAuthHandler) Login(c *fiber.Ctx) error {
	var req schemas.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	user, err := h.authService.Login(req.Username, req.Password)
	if err != nil {
		return err
	}

	fmt.Println(user)

	return nil
}

func (h *FiberAuthHandler) Register(c *fiber.Ctx) error {
	var req schemas.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	user, err := entities.NewUser(
		0,
		req.Name,
		req.Email,
		req.Username,
		req.Password,
	)

	if err != nil {
		return err
	}

	err = h.authService.Register(user)
	if err != nil {
		return err
	}

	fmt.Println(user)

	return nil
}
