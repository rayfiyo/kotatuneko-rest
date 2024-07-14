package user

import (
	"context"

	entity "github.com/rayfiyo/kotatuneko-rest/gen/user/resources"
)

type IRepository interface {
	Create(ctx context.Context, userName string) (*entity.User, error)
	SelectByID(ctx context.Context, userId string) (*entity.User, error)
	// Update(ctx context.Context, user *entity.User) error
	// Delete(ctx context.Context, userId *string) error
}
