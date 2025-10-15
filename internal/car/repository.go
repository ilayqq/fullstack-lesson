package car

import (
	"fullstack/config"
	"fullstack/domain"
)

type Repository interface {
	FindAll() ([]*domain.Car, error)
	GetByID(id int) (car *domain.Car, err error)
	Create(*domain.Car) error
	Update(id int, car *domain.Car) error
	Delete(id int) error
}

type RepositoryImpl struct{}

func NewRepository() Repository { return &RepositoryImpl{} }

func (r *RepositoryImpl) FindAll() ([]*domain.Car, error) {
	var car []*domain.Car
	err := config.DB.Find(&car).Error
	return car, err
}

func (r *RepositoryImpl) GetByID(id int) (car *domain.Car, err error) {
	err = config.DB.First(&car, id).Error
	return car, err
}

func (r *RepositoryImpl) Create(car *domain.Car) error {
	return config.DB.Create(car).Error
}

func (r *RepositoryImpl) Update(id int, car *domain.Car) error {
	car.ID = id
	return config.DB.Save(car).Error
}

func (r *RepositoryImpl) Delete(id int) error {
	return config.DB.Delete(id).Error
}
