package domain

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:255"`
	Email    string `gorm:"size:255;unique"`
	Password string `gorm:"size:255"`
}
