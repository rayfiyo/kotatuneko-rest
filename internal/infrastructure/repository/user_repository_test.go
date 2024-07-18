package repository

import (
	"context"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/rayfiyo/kotatuneko-rest/internal/domain/entity/user"
	"github.com/rayfiyo/kotatuneko-rest/internal/infrastructure/db"
	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) *sqlx.DB {
	db := db.New()
	_, err := db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id VARCHAR(36) PRIMARY KEY,
            name VARCHAR(50) UNIQUE NOT NULL,
            password VARCHAR(60) NOT NULL,
            score INT NOT NULL DEFAULT 0
        )
    `)
	if err != nil {
		t.Fatalf("Failed to create table: %v", err)
	}
	return db
}

func teardownTestDB(db *sqlx.DB) {
	db.Exec(`DROP TABLE IF EXISTS users`)
	db.Close()
}

func TestUserRepository(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(db)

	repo := NewUserRepository(db)

	ctx := context.TODO()
	user := &entity.User{
		Id:       "test-id",
		Name:     "test-user",
		Password: "test-password",
		Score:    100,
	}

	// Create のテスト
	err := repo.Create(ctx, user)
	assert.NoError(t, err)

	// SelectByName のテスト
	fetchedUser, err := repo.SelectByName(ctx, "test-user")
	assert.NoError(t, err)
	assert.NotNil(t, fetchedUser)
	assert.Equal(t, user.Name, fetchedUser.Name)

	// Update のテスト
	user.Score = 200
	err = repo.Update(ctx, user)
	assert.NoError(t, err)

	// GetRanking のテスト
	users, err := repo.GetRanking(ctx, 10)
	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, user.Score, users[0].Score)

	// Delete のテスト
	err = repo.Delete(ctx, user.Id)
	assert.NoError(t, err)

	// Confirm Deletion
	fetchedNilUser, err := repo.SelectByName(ctx, "test-user")
	assert.NoError(t, err)
	assert.Nil(t, fetchedNilUser)
}
