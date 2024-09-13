package services

import (
	"github.com/be2bag/myapp/internal/core/domain" // นำเข้าโมเดล User จาก domain
	"github.com/be2bag/myapp/internal/core/ports"  // นำเข้า interfaces ของ service และ repository จาก ports
)

type loginService struct {
	loginRepo ports.LoginRepository
}

func NewLoginService(loginRepo ports.LoginRepository) ports.LoginService {
	return &loginService{loginRepo: loginRepo}
}

func (s *loginService) LoginUser(user domain.User) (domain.User, error) {
	return s.loginRepo.CheckUser(user)
}
