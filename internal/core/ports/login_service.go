package ports

import "github.com/be2bag/myapp/internal/core/domain"

type LoginService interface {
	LoginUser(user domain.User) (domain.User, error)
}
