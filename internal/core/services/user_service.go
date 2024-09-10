package services

import (
	"github.com/be2bag/myapp/internal/core/domain"
	"github.com/be2bag/myapp/internal/core/ports"
)

type userService struct {
	userRepo ports.UserRepository
}

func NewUserService(userRepo ports.UserRepository) ports.UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) RegisterUser(user domain.User) (domain.User, error) {
	return s.userRepo.CreateUser(user)
}

func (s *userService) RemoveUser(id uint) error {
	return s.userRepo.DeleteUser(id)
}

func (s *userService) ModifyUser(user domain.User) (domain.User, error) {
	return s.userRepo.UpdateUser(user)
}

func (s *userService) FindUserByID(id uint) (domain.User, error) {
	return s.userRepo.GetUserByID(id)
}

func (s *userService) ListUsers() ([]domain.User, error) {
	return s.userRepo.GetAllUsers()
}
