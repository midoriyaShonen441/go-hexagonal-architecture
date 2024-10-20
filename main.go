package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/midoriyaShonen441/hexagonal-architecture/internal/adapters/handlers"
	"github.com/midoriyaShonen441/hexagonal-architecture/internal/adapters/repositories"
	"github.com/midoriyaShonen441/hexagonal-architecture/internal/adapters/repositories/models"
	"github.com/midoriyaShonen441/hexagonal-architecture/internal/config"
	"github.com/midoriyaShonen441/hexagonal-architecture/internal/core/services"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()

	db, err := gorm.Open(sqlite.Open(cfg.DBPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(&models.UserModel{})

	authRepo := repositories.NewGormAuthRepository(db)
	authService := services.NewAuthService(authRepo)
	authHandler := handlers.NewFiberAuthHandler(authService)

	app := fiber.New()
	v1 := app.Group("/api/v1")

	authRoutes := v1.Group("/auth")
	authRoutes.Post("/login", authHandler.Login)
	authRoutes.Post("/register", authHandler.Register)

	log.Fatal(app.Listen(":8080"))
}
