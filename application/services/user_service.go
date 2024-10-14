package services

import (
	"github.com/juliofilizzola/bot_discord/application/domain/model"
	"github.com/juliofilizzola/bot_discord/application/domain/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.repo.CreateUser(user)
}

func (s *UserService) GetUserByID(id string) (*model.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *UserService) UpdateUser(user *model.User) error {
	return s.repo.UpdateUser(user)
}

func (s *UserService) DeleteUser(id string) error {
	return s.repo.DeleteUser(id)
}
