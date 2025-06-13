package repositories

import (
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/adapters/persistence/models"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entities"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) Create(user *entities.User) error {
	userModel := &models.User{}
	userModel.FromEntity(user)

	if err := r.db.Create(userModel).Error; err != nil {
		return err
	}

	*user = *userModel.ToEntity()
	return nil
}
