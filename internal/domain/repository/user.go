package user

import (
	"context"

	"github.com/rayfiyo/kotatuneko-rest/proto"
)

type IRepository interface {
	Create(ctx context.Context, user *proto.User) error
	GetByID(ctx context.Context, id string) (*proto.User, error)
	Update(ctx context.Context, user *proto.User) error
	Delete(ctx context.Context, id *string) error
}
