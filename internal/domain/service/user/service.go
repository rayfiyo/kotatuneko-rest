package service

import (
	"context"
	"log"

	repository "github.com/rayfiyo/kotatuneko-rest/internal/domain/repository/user"
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
func (us *UserService) HashPassword(password string) (string, error) {
	// bcrypt.GenerateFromPassword の制約（73byte以上はだめ）
	if len([]byte(password)) > 72 {
		log.Printf("Password is too long: %v byte", len([]byte(password)))
		return "", errors.ErrInvalidUsernameOrPassword
	}

	// ハッシュ化
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Hashing: %v", err)
		return "", err
	}
	hashed := string(hashedByte)

	// 検証
	if err := bcrypt.CompareHashAndPassword(
		[]byte(hashed), []byte(password),
	); err != nil {
		log.Printf("Password verification fails (internal hashing): %v", err)
		return "", err
	}

	return hashed, nil
}

// ユーザーの認証のロジック
func (us *UserService) VerifyUserCredentials(ctx context.Context, userName string, password string) error {
	user, err := us.userRepository.SelectByName(ctx, userName)
	if err != nil {
		log.Printf("Selecting user by user name: %v", err)
		return err
	}

	if err := bcrypt.CompareHashAndPassword(
		// ハッシュ済み, 未ハッシュ
		[]byte(user.Password), []byte(password),
	); err != nil {
		log.Printf("Username or/and Password do not match: %v", err)
		return errors.ErrInvalidUsernameOrPassword
	}
	return nil
}
