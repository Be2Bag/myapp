package handlers

import (
	"strconv"

	"github.com/be2bag/myapp/internal/core/domain" // Import โครงสร้างข้อมูลของ User จาก domain
	"github.com/be2bag/myapp/internal/core/ports"  // Import interface ของบริการ UserService
	"golang.org/x/crypto/bcrypt"

	"github.com/gofiber/fiber/v2" // Import Fiber สำหรับสร้าง web framework ใน Go
)

// Struct สำหรับ handler ที่เชื่อมต่อกับบริการ userService
type userHandler struct {
	userService ports.UserService // userService คือ service interface ที่ใช้จัดการการทำงานของผู้ใช้
}

// ฟังก์ชันสร้าง userHandler ใหม่และลงทะเบียนเส้นทาง (routes) ของ HTTP requests
func NewUserHandler(app *fiber.App, userService ports.UserService) {
	handler := &userHandler{userService: userService} // สร้าง userHandler ใหม่พร้อมกับเชื่อมโยงกับ userService

	// ลงทะเบียนเส้นทางต่างๆ ของ HTTP requests และชี้ไปที่ method ของ handler ที่จะทำงาน
	app.Post("/users", handler.RegisterUser)  // ลงทะเบียนผู้ใช้ใหม่
	app.Delete("/delete", handler.RemoveUser) // ลบผู้ใช้โดยระบุ ID
	app.Put("/update", handler.ModifyUser)    // แก้ไขข้อมูลผู้ใช้โดยระบุ ID
	app.Get("/usersid", handler.FindUserByID) // ค้นหาผู้ใช้โดยระบุ ID
	app.Get("/users", handler.ListUsers)      // ดึงรายชื่อผู้ใช้ทั้งหมด
}

// ฟังก์ชันสำหรับจัดการ HTTP POST request เพื่อสมัครสมาชิกใหม่
func (h *userHandler) RegisterUser(c *fiber.Ctx) error {
	var user domain.User                        // สร้างตัวแปร user เพื่อรับข้อมูลผู้ใช้
	if err := c.BodyParser(&user); err != nil { // อ่านข้อมูลจาก body ของ request แล้วแปลงเป็น user struct
		return c.Status(fiber.StatusBadRequest).SendString(err.Error()) // ถ้าอ่านข้อมูลไม่สำเร็จ ตอบกลับด้วยข้อผิดพลาด 400 (Bad Request)
	}

	// เข้ารหัสรหัสผ่านด้วย bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to hash password")
	}
	user.Password = string(hashedPassword) // ตั้งค่ารหัสผ่านที่เข้ารหัสแล้วให้กับ user

	createdUser, err := h.userService.RegisterUser(user) // เรียกใช้บริการเพื่อสร้างผู้ใช้ใหม่
	if err != nil {                                      // ถ้ามีข้อผิดพลาดขณะสร้างผู้ใช้
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error()) // ตอบกลับด้วยข้อผิดพลาด 500 (Internal Server Error)
	}

	return c.Status(fiber.StatusCreated).JSON(createdUser) // ตอบกลับด้วยสถานะ 201 (Created) และข้อมูลผู้ใช้ที่ถูกสร้าง
}

// ฟังก์ชันสำหรับจัดการ HTTP DELETE request เพื่อทำการลบผู้ใช้ตาม ID
func (h *userHandler) RemoveUser(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Query("id"), 10, 32) // อ่านค่า ID จาก URL และแปลงเป็นจำนวนเต็ม (uint)
	err := h.userService.RemoveUser(uint(id))         // เรียกใช้บริการเพื่อทำการลบผู้ใช้ตาม ID
	if err != nil {                                   // ถ้าลบไม่สำเร็จ
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error()) // ตอบกลับด้วยข้อผิดพลาด 500 (Internal Server Error)
	}

	return c.SendStatus(fiber.StatusNoContent) // ตอบกลับด้วยสถานะ 204 (No Content) แสดงว่าลบสำเร็จแต่ไม่ต้องส่งข้อมูลกลับ
}

// ฟังก์ชันสำหรับจัดการ HTTP PUT request เพื่อทำการแก้ไขข้อมูลผู้ใช้ตาม ID
func (h *userHandler) ModifyUser(c *fiber.Ctx) error {
	var user domain.User                        // สร้างตัวแปร user เพื่อรับข้อมูลผู้ใช้ใหม่
	if err := c.BodyParser(&user); err != nil { // อ่านข้อมูลจาก body ของ request แล้วแปลงเป็น user struct
		return c.Status(fiber.StatusBadRequest).SendString(err.Error()) // ถ้าอ่านข้อมูลไม่สำเร็จ ตอบกลับด้วยข้อผิดพลาด 400 (Bad Request)
	}

	id, _ := strconv.ParseUint(c.Query("id"), 10, 32) // อ่านค่า ID จาก URL และแปลงเป็นจำนวนเต็ม (uint)
	user.ID = uint(id)                                // กำหนดค่า ID ของผู้ใช้ที่ต้องการแก้ไข

	updatedUser, err := h.userService.ModifyUser(user) // เรียกใช้บริการเพื่อทำการแก้ไขข้อมูลผู้ใช้
	if err != nil {                                    // ถ้าแก้ไขไม่สำเร็จ
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error()) // ตอบกลับด้วยข้อผิดพลาด 500 (Internal Server Error)
	}

	return c.JSON(updatedUser) // ตอบกลับด้วยข้อมูลผู้ใช้ที่ถูกแก้ไข
}

// ฟังก์ชันสำหรับจัดการ HTTP GET request เพื่อค้นหาข้อมูลผู้ใช้ตาม ID
func (h *userHandler) FindUserByID(c *fiber.Ctx) error {
	id, _ := strconv.ParseUint(c.Query("id"), 10, 32) // อ่านค่า ID จาก URL และแปลงเป็นจำนวนเต็ม (uint)
	user, err := h.userService.FindUserByID(uint(id)) // เรียกใช้บริการเพื่อค้นหาผู้ใช้ตาม ID
	if err != nil {                                   // ถ้าค้นหาไม่สำเร็จ (ไม่พบผู้ใช้)
		return c.Status(fiber.StatusNotFound).SendString(err.Error()) // ตอบกลับด้วยข้อผิดพลาด 404 (Not Found)
	}

	return c.JSON(user) // ตอบกลับด้วยข้อมูลผู้ใช้ที่พบ
}

// ฟังก์ชันสำหรับจัดการ HTTP GET request เพื่อดึงรายชื่อผู้ใช้ทั้งหมด
func (h *userHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.userService.ListUsers() // เรียกใช้บริการเพื่อดึงรายชื่อผู้ใช้ทั้งหมด
	if err != nil {                         // ถ้าดึงข้อมูลไม่สำเร็จ
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error()) // ตอบกลับด้วยข้อผิดพลาด 500 (Internal Server Error)
	}

	return c.JSON(users) // ตอบกลับด้วยข้อมูลรายชื่อผู้ใช้ทั้งหมดในรูปแบบ JSON
}
