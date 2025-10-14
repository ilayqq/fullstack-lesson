package car

import (
	"fullstack/config"
	"fullstack/domain"
)

type Repository interface {
	FindAll() ([]*domain.Car, error)
	Create(*domain.Car) error
	Update(*domain.Car) error
	Delete(*domain.Car) error
}

type RepositoryImpl struct{}

func NewRepository() Repository { return &RepositoryImpl{} }

func (r *RepositoryImpl) FindAll() ([]*domain.Car, error) {
	var car []*domain.Car
	err := config.DB.Find(&car).Error
	return car, err
}

func (r *RepositoryImpl) Create(car *domain.Car) error {
	return config.DB.Create(car).Error
}

func (r *RepositoryImpl) Update(car *domain.Car) error {
	return config.DB.Save(car).Error
}

func (r *RepositoryImpl) Delete(car *domain.Car) error {
	return config.DB.Delete(car).Error
}
