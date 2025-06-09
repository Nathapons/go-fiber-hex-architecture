package services

import (
	"errors"
	_ "errors"

	entities "github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/entites"
	"github.com/iamsamitdev/fiber-ecommerce-api/internal/core/domain/ports/repositories"
	"github.com/iamsamitdev/fiber-ecommerce-api/pkg/utils"
)

type AuthServiceImpl struct {
	userRepo repositories.UserRepository
}

func NewAuthServiceImpl(userRepo repositories.UserRepository) *AuthServiceImpl {
	return &AuthServiceImpl{
		userRepo: userRepo,
	}
}

func (s *AuthServiceImpl) Register(req entities.RegisterRequest) (*entities.User, error) {
	exitingUser, _ := s.userRepo.GetByEmail(req.Email)
	if exitingUser != nil {
		return nil, errors.New("User already exists")
	}

	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &entities.User{
		Email:     req.Email,
		Password:  hashPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      entities.RoleUser,
		IsActive:  true,
	}
	if err = s.userRepo.Create(user); err != nil {
		return nil, err
	}
	return user, err
}

func (s *AuthServiceImpl) Login(req entities.LoginRequest) (*entities.LoginResponse, error) {
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("Invalid login is not found")
	}

	if user.IsActive {
		return nil, errors.New("Account is deactive")
	}

	if !utils.CheckPassword(req.Password, user.Password) {
		return nil, errors.New("Invalid email or password")
	}

	token, err := utils.GenerateJWT(user.ID, string(user.Role))
	if err != nil {
		return nil, errors.New("Failed to generate token")
	}

	return &entities.LoginResponse{
		Token: token,
		User:  *user,
	}, nil

}

func (s *AuthServiceImpl) GetUserByID(id uint) (*entities.User, error) {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *AuthServiceImpl) UpdateUser(req entities.User) (*entities.User, error) {
	user, err := s.userRepo.GetByID(req.ID)
	if err != nil {
		return nil, err
	}

	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Email = req.Email

	if err := s.userRepo.Update(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *AuthServiceImpl) DeleteUser(id uint) error {
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return err
	}

	if err := s.userRepo.Delete(user.ID); err != nil {
		return err
	}
	return nil
}

func (s *AuthServiceImpl) GetAllUsers() ([]entities.User, error) {
	users, err := s.userRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
