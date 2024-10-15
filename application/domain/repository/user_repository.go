package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/juliofilizzola/bot_discord/application/domain/model"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByID(id string) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id string) error
	GetUserByGithubUsername(username string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetUserByID(id string) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) DeleteUser(id string) error {
	return r.db.Delete(&model.User{}, id).Error
}

func (r *userRepository) GetUserByGithubUsername(username string) (*model.User, error) {
	var user model.User
	if err := r.db.Raw("SELECT * FROM users WHERE github_username = ?", username).Scan(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to find user by GitHub username %s: %w", username, err)
	}
	return &user, nil

}
