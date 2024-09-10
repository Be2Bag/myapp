package ports

import "github.com/be2bag/myapp/internal/core/domain"

type UserService interface {
	RegisterUser(user domain.User) (domain.User, error)
	RemoveUser(id uint) error
	ModifyUser(user domain.User) (domain.User, error)
	FindUserByID(id uint) (domain.User, error)
	ListUsers() ([]domain.User, error)
}
