package services

import entities "github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entites"

type AuthService interface {
	Register(request *entities.RegisterRequest) (*entities.LoginResponse, error)
	Login(request *entities.LoginRequest) (*entities.LoginResponse, error)
	GetByUserID(id uint) (*entities.User, error)
	UpdateUser(user *entities.User) error
}
