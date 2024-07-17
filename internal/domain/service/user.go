package service

import (
	"context"

	"github.com/rayfiyo/kotatuneko-rest/gen/user/entity"
)

type User struct {
	entityUser  *entity.User
	userService *UserService
}

func NewUser(user *entity.User, userService *UserService) *User {
	return &User{
		entityUser:  user,
		userService: userService,
	}
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
func (u *User) Authenticate(ctx context.Context) error {
	return u.userService.VerifyUserCredentials(ctx, u.entityUser.Id, u.entityUser.Password)
}
