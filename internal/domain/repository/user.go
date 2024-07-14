package user

import (
	"github.com/rayfiyo/kotatuneko-rest/proto"
)

type IRepository interface {
	Create(user *proto.User) error
	GetByID(id int) (*proto.User, error)
	Update(user *proto.User) error
	Delete(id int) error
}
