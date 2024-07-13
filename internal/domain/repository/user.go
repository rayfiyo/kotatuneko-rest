package user

import (
	"github.com/rayfiyo/kotatuneko-rest/internal/domain/entity"
)

type IRepository interface {
	Create(user *entity.User) error
	GetByID(id int) (*entity.User, error)
	Update(user *entity.User) error
	Delete(id int) error
}
