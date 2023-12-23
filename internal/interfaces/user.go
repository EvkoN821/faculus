package interfaces

import (
	"github.com/IlyaZayats/faculus/internal/entity"
)

type UserRepository interface {
	Login(user entity.User) (int, error)
}
