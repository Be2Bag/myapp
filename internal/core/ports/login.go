package ports

import "github.com/be2bag/myapp/internal/core/domain"

type LoginRepository interface {
	CheckUser(user domain.User) (domain.User, error)
}
