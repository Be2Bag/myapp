package repositories

import (
	"github.com/be2bag/myapp/internal/core/domain"
	"github.com/be2bag/myapp/internal/core/ports"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) ports.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user domain.User) (domain.User, error) {
	result := r.db.Create(&user)
	return user, result.Error
}

func (r *userRepository) DeleteUser(id uint) error {
	result := r.db.Delete(&domain.User{}, id)
	return result.Error
}

func (r *userRepository) UpdateUser(user domain.User) (domain.User, error) {
	result := r.db.Save(&user)
	return user, result.Error
}

func (r *userRepository) GetUserByID(id uint) (domain.User, error) {
	var user domain.User
	result := r.db.First(&user, id)
	return user, result.Error
}

func (r *userRepository) GetAllUsers() ([]domain.User, error) {
	var users []domain.User
	result := r.db.Find(&users)
	return users, result.Error
}
