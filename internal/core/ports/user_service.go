package ports

import "github.com/be2bag/myapp/internal/core/domain"

// UserService เป็นอินเทอร์เฟซที่กำหนดฟังก์ชันหลักๆ ที่เกี่ยวข้องกับการจัดการผู้ใช้
type UserService interface {
	RegisterUser(user domain.User) (domain.User, error) // ฟังก์ชันสำหรับลงทะเบียนผู้ใช้ใหม่
	RemoveUser(id uint) error                           // ฟังก์ชันสำหรับลบผู้ใช้ตาม ID
	ModifyUser(user domain.User) (domain.User, error)   // ฟังก์ชันสำหรับแก้ไขข้อมูลผู้ใช้
	FindUserByID(id uint) (domain.User, error)          // ฟังก์ชันสำหรับค้นหาผู้ใช้ตาม ID
	ListUsers() ([]domain.User, error)                  // ฟังก์ชันสำหรับดึงรายการผู้ใช้ทั้งหมด
}
