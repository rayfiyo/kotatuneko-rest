package domain_repository

import (
	"context"

	"github.com/rayfiyo/kotatuneko-rest/internal/domain/entity/user"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	SelectByName(ctx context.Context, userName string) (*entity.User, error)
	// SelectByID(ctx context.Context, userId uuid.UUID) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id string) error
	GetRanking(ctx context.Context, limit int) ([]*entity.User, error) // limit は人数
}
