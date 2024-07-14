package repository

import (
	entity "github.com/rayfiyo/kotatuneko-rest/gen/user/resources"
)

type UserRepository interface {
	CreateUser(userName string) (*entity.User, error)
	SelectUserByID(userId string) (*entity.User, error)
	UpdateUser(user *entity.User) error
	DeleteUser(userId string) error
}
