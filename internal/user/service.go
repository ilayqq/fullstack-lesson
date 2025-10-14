package user

import "fullstack/domain"

type Service interface {
	GetById(id int) (*domain.User, error)
}
type ServiceImpl struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &ServiceImpl{repository}
}

func (s *ServiceImpl) GetById(id int) (*domain.User, error) {
	return s.repository.GetById(id)
}
