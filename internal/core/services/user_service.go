package services

import (
	"fmt"

	"github.com/be2bag/myapp/internal/core/domain" // นำเข้าโมเดล User จาก domain
	"github.com/be2bag/myapp/internal/core/ports"  // นำเข้า interfaces ของ service และ repository จาก ports
)

// userService เป็นโครงสร้างที่ทำงานเป็น service ของผู้ใช้ โดยใช้ repository ในการจัดการข้อมูล
type userService struct {
	userRepo ports.UserRepository // ใช้ userRepo สำหรับการจัดการฐานข้อมูลของผู้ใช้
}

// NewUserService เป็นฟังก์ชันที่สร้าง instance ของ userService และ return เป็น UserService interface
func NewUserService(userRepo ports.UserRepository) ports.UserService {
	return &userService{userRepo: userRepo}
}

// RegisterUser เรียกใช้ฟังก์ชัน CreateUser ของ repository เพื่อลงทะเบียนผู้ใช้ใหม่
func (s *userService) RegisterUser(user domain.User) (domain.User, error) {
	return s.userRepo.CreateUser(user)
}

// RemoveUser เรียกใช้ฟังก์ชัน DeleteUser ของ repository เพื่อลบผู้ใช้ตาม ID
func (s *userService) RemoveUser(id uint) error {
	return s.userRepo.DeleteUser(id)
}

// ModifyUser เรียกใช้ฟังก์ชัน UpdateUser ของ repository เพื่อแก้ไขข้อมูลผู้ใช้
func (s *userService) ModifyUser(user domain.User) (domain.User, error) {
	return s.userRepo.UpdateUser(user)
}

// FindUserByID เรียกใช้ฟังก์ชัน GetUserByID ของ repository เพื่อนำข้อมูลผู้ใช้จากฐานข้อมูลตาม ID
func (s *userService) FindUserByID(id uint) (domain.User, error) {
	fmt.Println("FindUserByID")
	return s.userRepo.GetUserByID(id)
}

// ListUsers เรียกใช้ฟังก์ชัน GetAllUsers ของ repository เพื่อดึงข้อมูลผู้ใช้ทั้งหมด
func (s *userService) ListUsers() ([]domain.User, error) {
	fmt.Println("ListUsers")
	return s.userRepo.GetAllUsers()
}
