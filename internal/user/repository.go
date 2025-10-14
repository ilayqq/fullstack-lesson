package user

import (
	"fullstack/config"
	"fullstack/domain"
)

type Repository interface {
	GetById(userID int) (*domain.User, error)
}

type RepositoryImpl struct{}

func NewRepository() Repository {
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) GetById(userID int) (*domain.User, error) {
	var user domain.User
	result := config.DB.Preload("Cars").First(&user, userID)
	return &user, result.Error
}
