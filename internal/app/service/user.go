package service

import (
	"context"

	"github.com/rayfiyo/kotatuneko-rest/internal/domain/repository"
	"github.com/rayfiyo/kotatuneko-rest/proto"
)

type IUserService interface {
	CreateByID(ctx context.Context, userID string) error
	GetByID(ctx context.Context, userID string) (*proto.User, error)
	// Update(ctx context.Context, userID string, input *schema.UpdateUserInput) error
}

func New(repo user.IRepository) {
}
