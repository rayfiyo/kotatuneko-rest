package service

import (

	// "github.com/google/uuid"
	entity "github.com/rayfiyo/kotatuneko-rest/gen/user/resources"
	// schema "github.com/rayfiyo/kotatuneko-rest/gen/user/rpc"
	// "github.com/rayfiyo/kotatuneko-rest/internal/domain/repository"
)

type User struct {
	entityUser  *entity.User
	userService *UserService
}

func NewUser(user *entity.User) *User {
	return &User{entityUser: user}
}

// ハッシュ化
func (u *User) Hashing() error {
	hashed, err := u.userService.HashPassword(u.entityUser.Password)
	if err != nil {
		return err
	}
	u.entityUser.Password = hashed
	return nil
}

// 認証
func (u *User) Authenticate() (bool, error) {
	return u.userService.VerifyUserCredentials(u.entityUser.Id, u.entityUser.Password)
}
