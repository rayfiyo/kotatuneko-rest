package service

import (
	"log"

	"github.com/google/uuid"
	entity "github.com/rayfiyo/kotatuneko-rest/gen/user/resources"
	schema "github.com/rayfiyo/kotatuneko-rest/gen/user/rpc"
	repository "github.com/rayfiyo/kotatuneko-rest/internal/domain/user"
)

// 依存性逆転の原則
type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepo}
}

func (s *UserService) CreateUser(user *entity.User, password string) (*entity.User, error) {
	user, err := s.userRepository.Create(user, password)

	// UUID の生成
	user.Id = uuid.New().String()

	if err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUserByID(userID string) (*entity.User, error) {
	user, err := s.userRepository.SelectByID(userID)
	if err != nil {
		log.Printf("Error getting user by ID: %v", err)
		return nil, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(req *schema.UpdateUserRequest) error {
	user := &entity.User{
		Id:   req.User.Id,
		Name: req.User.Name,
	}
	if err := s.userRepository.UpdateUser(user); err != nil {
		log.Printf("Error updating user: %v", err)
		return err
	}
	return nil
}

func (s *UserService) DeleteUser(userID string) error {
	if err := s.userRepository.DeleteUser(userID); err != nil {
		log.Printf("Error deleting user: %v", err)
		return err
	}
	return nil
}
