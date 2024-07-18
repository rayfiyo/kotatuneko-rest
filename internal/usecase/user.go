package usecase

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/rayfiyo/kotatuneko-rest/internal/domain/entity/user"
	repository "github.com/rayfiyo/kotatuneko-rest/internal/domain/repository/user"
	"github.com/rayfiyo/kotatuneko-rest/internal/domain/service/user"
	"github.com/rayfiyo/kotatuneko-rest/pkg/errors"
)

type UserUsecase interface {
	RegisterUser(ctx context.Context, userName, password string) (*entity.User, *errors.Error)
	LoginUser(ctx context.Context, userName, password string) (*entity.User, error)
	// LogoutUser(ctx context.Context, userID string) error
	UpdateUser(ctx context.Context, user *entity.User) error
	GetRanking(ctx context.Context, limit int) ([]*entity.User, error) // リミットは人数
}

type userUsecaseImpl struct {
	userRepo    repository.UserRepository
	userService *service.UserService
}

func NewUserUsecase(userRepo repository.UserRepository, userService *service.UserService) UserUsecase {
	return &userUsecaseImpl{
		userRepo:    userRepo,
		userService: userService,
	}
}

// ユーザー情報の作成
func (uc *userUsecaseImpl) RegisterUser(ctx context.Context, userName, password string) (*entity.User, *errors.Error) {
	// ユーザー名の重複チェック
	existingUser, err := uc.userRepo.SelectByName(ctx, userName)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[warning]行が無いけど，続行します: %v", err)
		} else {
			log.Printf("Error checking duplicate user at registration: %v", err)
			newErr := errors.New(http.StatusBadRequest, err)
			return nil, newErr
		}
	}
	if existingUser != nil {
		return nil, errors.ErrUserAlreadyExists
	}

	// パスワードのハッシュ化
	if password == "" {
		return nil, errors.ErrInvalidUsernameOrPassword
	}
	hashedPassword, err := uc.userService.HashPassword(password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return nil, err
	}

	// uuid の生成
	uuid := uuid.New().String()

	// ハイフン削除
	// uuidWithoutHyphen := strings.Replace(uuid, "-", "", -1)

	// ユーザーの作成
	user := &entity.User{
		Id:       uuid,
		Name:     userName,
		Password: hashedPassword,
		Score:    0,
	}
	if err := uc.userRepo.Create(ctx, user); err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}

	return user, nil
}

func (uc *userUsecaseImpl) LoginUser(ctx context.Context, userName, password string) (*entity.User, error) {
	if err := uc.userService.VerifyUserCredentials(ctx, userName, password); err != nil {
		log.Printf("Error verifying user credentials: %v", err)
		return nil, errors.ErrInvalidUsernameOrPassword
	}

	user, err := uc.userRepo.SelectByName(ctx, userName)
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		return nil, err
	}

	return user, nil
}

// ユーザー情報の更新
func (uc *userUsecaseImpl) UpdateUser(ctx context.Context, user *entity.User) error {
	existingUser, err := uc.userRepo.SelectByName(ctx, user.Name)
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		return err
	}
	if existingUser == nil {
		return errors.ErrUserNotFound
	}

	// パスワードのハッシュ化
	if user.Password == "" {
		return errors.ErrInvalidUsernameOrPassword
	}
	hashedPassword, err := uc.userService.HashPassword(user.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		return err
	}
	user.Password = hashedPassword

	if err := uc.userRepo.Update(ctx, user); err != nil {
		log.Printf("Error updating user: %v", err)
		return err
	}

	return nil
}

// ランキングの取得
func (uc *userUsecaseImpl) GetRanking(ctx context.Context, limit int) ([]*entity.User, error) {
	users, err := uc.userRepo.GetRanking(ctx, limit)
	if err != nil {
		log.Printf("Error fetching top users by score: %v", err)
		return nil, err
	}
	return users, nil
}
