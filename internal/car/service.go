package car

import (
	"fullstack/domain"
)

type Service interface {
	GetAllCars() ([]*domain.Car, error)
	Create(car *domain.Car) error
	Update(id int, car *domain.Car) (*domain.Car, error)
	Delete(id int) error
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

func (s *ServiceImpl) Update(id int, car *domain.Car) (*domain.Car, error) {
	if err := s.repository.Update(id, car); err != nil {
		return nil, err
	}

	updatedCar, err := s.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return updatedCar, nil
}

func (s *ServiceImpl) Delete(id int) error {
	return s.repository.Delete(id)
}
