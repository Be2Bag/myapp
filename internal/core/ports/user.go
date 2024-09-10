package ports

import "github.com/be2bag/myapp/internal/core/domain"

type UserRepository interface {
	CreateUser(user domain.User) (domain.User, error)
	DeleteUser(id uint) error
	UpdateUser(user domain.User) (domain.User, error)
	GetUserByID(id uint) (domain.User, error)
	GetAllUsers() ([]domain.User, error)
}
