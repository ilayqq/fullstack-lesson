package main

import (
	"fullstack/config"
	"fullstack/internal/car"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	config.InitDB()

	carRepository := &car.RepositoryImpl{}
	carService := car.NewCar(carRepository)

	carService.GetAllCars()
}
