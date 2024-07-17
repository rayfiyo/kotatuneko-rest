package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rayfiyo/kotatuneko-rest/pkg/config"
	"github.com/rayfiyo/kotatuneko-rest/pkg/errors"
)

// データベース接続の設定
func Init() (*sqlx.DB, error) {
	if err := config.LoadEnv(); err != nil {
		log.Printf("Failed to open .env: %v", err)
		return nil, err
	}

	if config.Mode == "dev" {
		db, err := sqlx.Open("postgres", config.DB_DSN)
		if err != nil {
			log.Printf("Failed to open DB: %v", err)
			return nil, err
		}
		log.Println(db)
		return db, nil
	}
	return nil, errors.New("mode is invalid: " + config.Mode)
}

func New() *sqlx.DB {
	db, err := Init()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}
