package repositories

import (
	"github.com/be2bag/myapp/internal/core/domain" // นำเข้าโมเดล User จาก domain
	"github.com/be2bag/myapp/internal/core/ports"  // นำเข้า interfaces ของ repository จาก ports

	"gorm.io/gorm" // นำเข้า GORM สำหรับจัดการฐานข้อมูล
)

type loginRepository struct {
	db *gorm.DB
}

func NewLoginRepository(db *gorm.DB) ports.LoginRepository {
	return &loginRepository{db: db}
}

func (r *loginRepository) CheckUser(user domain.User) (domain.User, error) {
	var users domain.User
	result := r.db.Where("email = ?", user.Email).First(&users)
	return users, result.Error
}
