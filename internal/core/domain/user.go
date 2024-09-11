package domain

type User struct {
	ID       uint   `gorm:"primaryKey"`      // ฟิลด์ ID กำหนดเป็นคีย์หลัก (primary key) ของตารางในฐานข้อมูล
	Name     string `gorm:"size:255"`        // ฟิลด์ Name เป็นสตริง มีขนาดไม่เกิน 255 ตัวอักษร
	Email    string `gorm:"size:255;unique"` // ฟิลด์ Email เป็นสตริง ขนาดไม่เกิน 255 ตัวอักษร และต้องไม่ซ้ำกัน (unique)
	Password string `gorm:"size:255"`        // ฟิลด์ Password เป็นสตริง ขนาดไม่เกิน 255 ตัวอักษร สำหรับเก็บรหัสผ่าน
}
