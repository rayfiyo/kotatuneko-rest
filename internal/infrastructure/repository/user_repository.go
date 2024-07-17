package repository

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/rayfiyo/kotatuneko-rest/gen/user/entity"
	"github.com/rayfiyo/kotatuneko-rest/internal/domain/repository/user"
)

type userRepositoryImpl struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) domain_repository.UserRepository {
	return &userRepositoryImpl{db: db}
}

// ユーザーを作成するDBの操作
func (r *userRepositoryImpl) Create(ctx context.Context, user *entity.User) error {
	query := `INSERT INTO users (id, name, password, score) VALUES (:id, :name, :password, :score)`
	_, err := r.db.NamedExecContext(ctx, query, user)
	return err
}

// user_name からユーザー（entity）を持ってくる
// 存在しないときのハンドリングすること
func (r *userRepositoryImpl) SelectByName(ctx context.Context, userName string) (*entity.User, error) {
	var user entity.User
	query := `SELECT id, name, password, score FROM users WHERE name = $1`
	err := r.db.GetContext(ctx, &user, query, userName)
	if err == sql.ErrNoRows { // 条件分岐する意味あんまないけど，明示
		return nil, err
	}
	return &user, err
}

// ユーザーの情報を更新
func (r *userRepositoryImpl) Update(ctx context.Context, user *entity.User) error {
	query := `UPDATE users SET name = :name, password = :password, score = :score WHERE id = :id`
	_, err := r.db.NamedExecContext(ctx, query, user)

	return err
}

// ユーザーの情報を削除
func (r *userRepositoryImpl) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

// ランキングを取得
func (r *userRepositoryImpl) GetRanking(ctx context.Context, limit int) ([]*entity.User, error) {
	var users []*entity.User

	query := `SELECT id, name, score FROM users ORDER BY score DESC LIMIT $1`
	err := r.db.SelectContext(ctx, &users, query, limit)
	return users, err
}
