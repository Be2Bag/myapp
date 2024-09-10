package ports

import "github.com/be2bag/myapp/internal/core/domain"

// UserRepository เป็นอินเทอร์เฟซที่กำหนดวิธีการจัดการข้อมูลของผู้ใช้ในระดับฐานข้อมูล
type UserRepository interface {
	CreateUser(user domain.User) (domain.User, error) // ฟังก์ชันสำหรับสร้างผู้ใช้ใหม่ในฐานข้อมูล
	DeleteUser(id uint) error                         // ฟังก์ชันสำหรับลบผู้ใช้จากฐานข้อมูลตาม ID
	UpdateUser(user domain.User) (domain.User, error) // ฟังก์ชันสำหรับอัปเดตข้อมูลผู้ใช้ในฐานข้อมูล
	GetUserByID(id uint) (domain.User, error)         // ฟังก์ชันสำหรับดึงข้อมูลผู้ใช้จากฐานข้อมูลตาม ID
	GetAllUsers() ([]domain.User, error)              // ฟังก์ชันสำหรับดึงรายชื่อผู้ใช้ทั้งหมดจากฐานข้อมูล
}
