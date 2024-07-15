package database

import (
	"github.com/jmoiron/sqlx"

	entity "github.com/rayfiyo/kotatuneko-rest/gen/user/resources"
	// "github.com/rayfiyo/kotatuneko-rest/internal/domain/repository"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user *entity.User) error {
	query := "INSERT INTO users (user_id, user_name) VALUES ($1, $2)"
	_, err := r.db.Exec(query, user.Id, user.Name)
	return err
}

func (r *UserRepository) GetUserByID(userID string) (*entity.User, error) {
	query := "SELECT user_id, user_name FROM users WHERE user_id = $1"
	row := r.db.QueryRow(query, userID)
	user := &entity.User{}
	err := row.Scan(&user.Id, &user.Name)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(user *entity.User) error {
	query := "UPDATE users SET user_name = $1 WHERE user_id = $2"
	_, err := r.db.Exec(query, user.Name, user.Id)
	return err
}

func (r *UserRepository) DeleteUser(userID string) error {
	query := "DELETE FROM users WHERE user_id = $1"
	_, err := r.db.Exec(query, userID)
	return err
}
