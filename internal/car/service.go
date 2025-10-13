package car

import "fmt"

type Service struct {
	repository Repository
}

func NewCar(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetAllCars() {
	cars := s.repository.FindAll()
	for _, car := range cars {
		fmt.Println(car)
	}
}
