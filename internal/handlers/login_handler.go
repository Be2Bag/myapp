package handlers

import (
	"os"
	"time"

	"github.com/be2bag/myapp/internal/core/domain" // Import โครงสร้างข้อมูลของ User จาก domain
	"github.com/be2bag/myapp/internal/core/ports"  // Import interface ของบริการ UserService
	"github.com/gofiber/fiber/v2"                  // Import Fiber สำหรับสร้าง web framework ใน Go
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type loginHandler struct {
	loginService ports.LoginService
}

var jwtSecret = []byte(os.Getenv("SECRETKEY"))

func NewLoginHandler(app *fiber.App, loginService ports.LoginService) {
	handlers := &loginHandler{loginService: loginService}

	app.Post("/login", handlers.LoginUser)

}

func (h *loginHandler) LoginUser(c *fiber.Ctx) error {
	var user domain.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}

	users, err := h.loginService.LoginUser(user)

	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}

	err = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(user.Password))

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid email or password")
	}

	// สร้าง JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": users.ID,
		"email":   users.Email,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // Token หมดอายุใน 72 ชั่วโมง
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not login")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"token":   tokenString,
	})
}
