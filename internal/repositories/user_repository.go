package repositories

import (
	"github.com/be2bag/myapp/internal/core/domain" // นำเข้าโมเดล User จาก domain
	"github.com/be2bag/myapp/internal/core/ports"  // นำเข้า interfaces ของ repository จาก ports

	"gorm.io/gorm" // นำเข้า GORM สำหรับจัดการฐานข้อมูล
)

// userRepository เป็นโครงสร้างที่เก็บการเชื่อมต่อฐานข้อมูลด้วย GORM
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository เป็นฟังก์ชันที่สร้าง instance ของ userRepository และ return เป็น UserRepository interface
func NewUserRepository(db *gorm.DB) ports.UserRepository {
	return &userRepository{db: db}
}

// CreateUser เป็นฟังก์ชันสำหรับเพิ่มผู้ใช้ใหม่ในฐานข้อมูล
func (r *userRepository) CreateUser(user domain.User) (domain.User, error) {
	result := r.db.Create(&user) // ใช้ GORM สร้างข้อมูลผู้ใช้ใหม่
	return user, result.Error    // คืนค่าผู้ใช้ที่ถูกสร้าง และ error ถ้ามี
}

// DeleteUser เป็นฟังก์ชันสำหรับลบผู้ใช้จากฐานข้อมูลตาม ID
func (r *userRepository) DeleteUser(id uint) error {
	result := r.db.Delete(&domain.User{}, id) // ใช้ GORM ลบข้อมูลผู้ใช้ตาม ID
	return result.Error                       // คืนค่า error ถ้ามี
}

// UpdateUser เป็นฟังก์ชันสำหรับอัปเดตข้อมูลผู้ใช้ในฐานข้อมูล
func (r *userRepository) UpdateUser(user domain.User) (domain.User, error) {
	result := r.db.Save(&user) // ใช้ GORM บันทึกการเปลี่ยนแปลงของข้อมูลผู้ใช้
	return user, result.Error  // คืนค่าผู้ใช้ที่ถูกอัปเดต และ error ถ้ามี
}

// GetUserByID เป็นฟังก์ชันสำหรับดึงข้อมูลผู้ใช้ตาม ID
func (r *userRepository) GetUserByID(id uint) (domain.User, error) {
	var user domain.User
	result := r.db.First(&user, id) // ใช้ GORM ดึงข้อมูลผู้ใช้ตาม ID
	return user, result.Error       // คืนค่าผู้ใช้ และ error ถ้ามี
}

// GetAllUsers เป็นฟังก์ชันสำหรับดึงข้อมูลผู้ใช้ทั้งหมด
func (r *userRepository) GetAllUsers() ([]domain.User, error) {
	var users []domain.User
	result := r.db.Find(&users) // ใช้ GORM ดึงข้อมูลผู้ใช้ทั้งหมด
	return users, result.Error  // คืนค่าผู้ใช้ทั้งหมด และ error ถ้ามี
}
