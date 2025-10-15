package main

import (
	"fullstack/config"
	"fullstack/internal/car"
	"github.com/gin-gonic/gin"
	"log"

	"github.com/joho/godotenv"

	_ "fullstack/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Full Stack
// @version 1.0
// @description API for fullstack leasson
// @host localhost:8080
func main() {
	r := gin.Default()

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	config.InitDB()

	carRepository := car.NewRepository()
	carService := car.NewService(carRepository)
	carHandler := car.NewHandler(carService)

	//userRepository := user.NewRepository()
	//userService := user.NewService(userRepository)

	api := r.Group("/api")
	{
		api.GET("/cars", carHandler.GetAllCars)
		api.POST("/cars", carHandler.CreateCar)
		api.PUT("/cars/:id", carHandler.UpdateCar)
		api.DELETE("/cars/:id", carHandler.DeleteCar)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
