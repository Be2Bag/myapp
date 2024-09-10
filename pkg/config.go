package pkg

import (
	"fmt" // ใช้ในการจัดการข้อความที่มีการแทรกค่าตัวแปร
	"log" // ใช้ในการแสดงข้อความ error หรือ log
	"os"  // ใช้ในการเข้าถึงตัวแปรสภาพแวดล้อม

	"github.com/joho/godotenv" // ใช้ในการโหลดไฟล์ .env เพื่อดึงค่าตัวแปรสภาพแวดล้อม
)

// Config เป็นโครงสร้างที่เก็บค่า URL ของฐานข้อมูล
type Config struct {
	DatabaseURL string // เก็บ URL สำหรับการเชื่อมต่อฐานข้อมูล MySQL
}

// LoadConfig ฟังก์ชันนี้ใช้ในการโหลดตัวแปรสภาพแวดล้อมจากไฟล์ .env และสร้าง URL สำหรับเชื่อมต่อกับฐานข้อมูล MySQL
func LoadConfig() *Config {
	// โหลดไฟล์ .env เพื่อดึงตัวแปรสภาพแวดล้อม
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file") // ถ้าเกิดข้อผิดพลาดในการโหลด .env จะทำการหยุดโปรแกรมและแสดงข้อความ error
	}

	// ดึงค่าตัวแปรจากไฟล์ .env
	host := os.Getenv("DB_HOST")         // ดึงค่า host ของฐานข้อมูล
	user := os.Getenv("DB_USER")         // ดึงค่า user ของฐานข้อมูล
	password := os.Getenv("DB_PASSWORD") // ดึงค่า password ของฐานข้อมูล
	port := os.Getenv("DB_PORT")         // ดึงค่า port ของฐานข้อมูล
	database := os.Getenv("DB_DATABASE") // ดึงค่า database ที่จะใช้งาน

	// สร้าง URL สำหรับการเชื่อมต่อ MySQL โดยแทรกค่าที่ดึงจาก .env เข้าไปในรูปแบบ URL
	databaseURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, database)

	// คืนค่าการตั้งค่าที่สร้างขึ้น
	return &Config{
		DatabaseURL: databaseURL, // เก็บ URL สำหรับใช้ในการเชื่อมต่อฐานข้อมูล
	}
}
