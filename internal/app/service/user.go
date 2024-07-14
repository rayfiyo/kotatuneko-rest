package service

import (
	entity "github.com/rayfiyo/kotatuneko-rest/gen/user/resources"
	schema "github.com/rayfiyo/kotatuneko-rest/gen/user/rpc"
	"github.com/rayfiyo/kotatuneko-rest/internal/domain/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepo}
}

func (s *UserService) CreateUser(userName string) (*entity.User, error) {
	user, err := s.userRepository.CreateUser(userName)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *UserService) GetUserByID(userId string) (*entity.User, error) {
	user, err := s.userRepository.SelectUserByID(userId)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(req *schema.UpdateUserRequest) error {
	user := &entity.User{
		Id:   req.User.Id,
		Name: req.User.Name,
	}
	if err := s.userRepository.UpdateUser(user); err != nil {
		return err
	}
	return nil
}

func (s *UserService) DeleteUser(userID string) error {
	if err := s.userRepository.DeleteUser(userID); err != nil {
		return err
	}
	return nil
}
