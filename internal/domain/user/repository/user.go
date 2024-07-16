package domain

import (
	"github.com/rayfiyo/kotatuneko-rest/gen/user/entity"
)

type UserRepository interface {
	Create(user *entity.User) *error
	SelectByName(userName string) (*entity.User, error)
	// SelectByID(userId uuid.UUID) (*entity.User, error)
	Update(user *entity.User) error
	Delete(id string) error
}
