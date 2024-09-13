package main

import (
	"github.com/be2bag/myapp/internal/core/domain"   // Import โครงสร้างของโมเดล User จาก domain
	"github.com/be2bag/myapp/internal/core/services" // Import บริการที่จัดการกับ UserService
	"github.com/be2bag/myapp/internal/handlers"      // Import Handlers สำหรับ HTTP request
	"github.com/be2bag/myapp/internal/repositories"  // Import repositories สำหรับเชื่อมต่อกับฐานข้อมูล
	"github.com/be2bag/myapp/pkg"                    // Import pkg สำหรับโหลดการตั้งค่า (config)

	"github.com/gofiber/fiber/v2" // Import Fiber สำหรับสร้าง web server
	"gorm.io/driver/mysql"        // Import MySQL driver สำหรับ GORM
	"gorm.io/gorm"                // Import GORM สำหรับจัดการฐานข้อมูล
)

func main() {
	app := fiber.New() // สร้างแอปพลิเคชัน Fiber ใหม่

	// โหลดการตั้งค่าจาก pkg (เช่น DatabaseURL)
	config := pkg.LoadConfig()

	// เชื่อมต่อกับฐานข้อมูล MySQL โดยใช้ GORM
	db, err := gorm.Open(mysql.Open(config.DatabaseURL), &gorm.Config{})
	if err != nil {
		// ถ้าเชื่อมต่อฐานข้อมูลไม่ได้ ให้แสดงข้อผิดพลาดและหยุดการทำงาน
		panic("failed to connect database")
	}

	// ทำการ Auto Migrate เพื่อตรวจสอบหรือสร้างตารางในฐานข้อมูลตามโมเดล User ที่กำหนด
	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		// ถ้าเกิดข้อผิดพลาดในการ Migrate ให้แสดงข้อผิดพลาดและหยุดการทำงาน
		panic("failed to migrate database: " + err.Error())
	}

	//======================= Login ============================
	loginRepo := repositories.NewLoginRepository(db)

	loginService := services.NewLoginService(loginRepo)

	handlers.NewLoginHandler(app, loginService)

	// app.Use(handlers.AuthMiddleware)

	// สร้าง Repository (คลาสที่จัดการกับฐานข้อมูล) สำหรับ User
	userRepo := repositories.NewUserRepository(db)

	// สร้าง Service (คลาสที่จัดการ logic ธุรกิจ) สำหรับ User
	userService := services.NewUserService(userRepo)

	// สร้าง Handlers (คลาสที่จัดการ HTTP request/response) และลงทะเบียน routes
	handlers.NewUserHandler(app, userService)

	// เริ่มต้นเซิร์ฟเวอร์ และฟัง request ที่พอร์ต 3000
	app.Listen(":3000")
}
