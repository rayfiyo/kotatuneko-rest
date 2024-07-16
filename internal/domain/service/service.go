package service

import (
	"context"
	"log"

	"github.com/rayfiyo/kotatuneko-rest/internal/domain/repository"
	"github.com/rayfiyo/kotatuneko-rest/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepo}
}

// パスワードのハッシュ化のロジック
func (us *UserService) HashPassword(password []byte) ([]byte, error) {
	// bcrypt.GenerateFromPassword の制約（73byte以上はだめ）
	if len(password) > 72 {
		log.Printf("Password is too long: %v", len([]byte(password)))
		return nil, errors.ErrInvalidUsernameOrPassword
	}

	// ハッシュ化
	hashed, err := bcrypt.GenerateFromPassword(
		[]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Hashing: %v", err)
		return nil, err
	}

	// 検証
	if err := bcrypt.CompareHashAndPassword(
		hashed, password,
	); err != nil {
		log.Printf("Password verification fails (internal hashing): %v", err)
		return nil, err
	}

	return hashed, nil
}

// ユーザーの認証のロジック
func (us *UserService) VerifyUserCredentials(ctx context.Context, userName string, password []byte) error {
	user, err := us.userRepository.SelectByName(ctx, userName)
	if err != nil {
		log.Printf("Selecting user by user name: %v", err)
		return err
	}

	if err := bcrypt.CompareHashAndPassword(
		// ハッシュ済み, 未ハッシュ
		user.Password, password,
	); err != nil {
		log.Printf("Username or/and Password do not match: %v", err)
		return err
	}
	return nil
}
