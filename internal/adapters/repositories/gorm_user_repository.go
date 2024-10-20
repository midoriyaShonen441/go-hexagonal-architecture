package repositories

import (
	"github.com/midoriyaShonen441/hexagonal-architecture/internal/adapters/repositories/models"
	"github.com/midoriyaShonen441/hexagonal-architecture/internal/core/domains/entities"
	"github.com/midoriyaShonen441/hexagonal-architecture/internal/core/ports"
	"gorm.io/gorm"
)

type gormAuthRepository struct {
	db *gorm.DB
}

func NewGormAuthRepository(db *gorm.DB) ports.AuthRepository {
	return &gormAuthRepository{db: db}
}

func (r *gormAuthRepository) Create(u *entities.User) error {
	userModel := models.NewUserModel(u)

	result := r.db.Create(userModel)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *gormAuthRepository) FindByUsername(username string) (*entities.User, error) {
	var userModel models.UserModel

	if err := r.db.Where("username = ?", username).First(&userModel).Error; err != nil {
		return nil, err
	}

	user, err := entities.NewUser(
		userModel.Id,
		userModel.Name,
		userModel.Email,
		userModel.Username,
		"Password12345",
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
