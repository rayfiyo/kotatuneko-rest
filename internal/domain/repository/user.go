package repository

import (
	"github.com/google/uuid"
	entity "github.com/rayfiyo/kotatuneko-rest/gen/user/resources"
)

type UserRepository interface {
	Create(user *entity.User) (*entity.User, error)
	SelectByID(userId uuid.UUID) (*entity.User, error)
	SelectByName(userName string) (*entity.User, error)
	// Update(user *entity.User) (*entity.User, error)
	Delete(user *entity.User) error
}
