package car

import (
	"fullstack/domain"
)

type Service interface {
	GetAllCars() ([]*domain.Car, error)
	Create(car *domain.Car) error
	Update(car *domain.Car) error
	Delete(car *domain.Car) error
}
type ServiceImpl struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &ServiceImpl{repository}
}

func (s *ServiceImpl) GetAllCars() ([]*domain.Car, error) {
	return s.repository.FindAll()
}

func (s *ServiceImpl) Create(car *domain.Car) error {
	return s.repository.Create(car)
}

func (s *ServiceImpl) Update(car *domain.Car) error {
	return s.repository.Update(car)
}

func (s *ServiceImpl) Delete(car *domain.Car) error {
	return s.repository.Delete(car)
}
