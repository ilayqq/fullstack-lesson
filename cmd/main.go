package main

import (
	"fullstack/config"
	"fullstack/domain"
	"fullstack/internal/car"
	"fullstack/internal/user"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	config.InitDB()

	carRepository := car.NewRepository()
	carService := car.NewService(carRepository)

	userRepository := user.NewRepository()
	userService := user.NewService(userRepository)

	cars, _ := carService.GetAllCars()
	for _, car := range cars {
		log.Println(car)
	}
	carService.Create(&domain.Car{Brand: "test", Color: "test", Model: "test", ModelYear: 123, Price: 123, RegistrationNumber: "test", UserID: 1})
	carService.Update(&domain.Car{ID: 1, Brand: "update", Color: "update", Model: "update", ModelYear: 123, Price: 123, RegistrationNumber: "update", UserID: 1})
	carService.Delete(&domain.Car{ID: 2})

	config.DB.Create(&domain.User{ID: 1, Name: "test"})

	userService.GetById(1)
}
