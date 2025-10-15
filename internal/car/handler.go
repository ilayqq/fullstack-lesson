package car

import (
	"fullstack/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type HandlerImpl struct {
	service Service
}

func NewHandler(service Service) *HandlerImpl { return &HandlerImpl{service} }

// GetAllCars godoc
// @Summary      Get all cars
// @Description  Returns a list of cars
// @Tags         cars
// @Accept       json
// @Produce      json
// @Success      200  {array}   domain.Car
// @Failure      500  {object}	map[string]string
// @Router       /api/cars [get]
func (h *HandlerImpl) GetAllCars(c *gin.Context) {
	cars, err := h.service.GetAllCars()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"cars": cars})
}

// CreateCar godoc
// @Summary      Create a new car
// @Description  Adds a new car to the database
// @Tags         cars
// @Accept       json
// @Produce      json
// @Param        car  body      domain.Car  true  "Car data"
// @Success      201  {object}  domain.Car
// @Failure      400  {object} 	map[string]string
// @Failure      500  {object} 	map[string]string
// @Router       /api/cars [post]
func (h *HandlerImpl) CreateCar(c *gin.Context) {
	var req *domain.Car
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	car := h.service.Create(req)

	c.JSON(http.StatusCreated, gin.H{"car": car})
}

// UpdateCar godoc
// @Summary      Update a car
// @Description  Updates an existing car by ID
// @Tags         cars
// @Accept       json
// @Produce      json
// @Param        id   path      int                    true  "Car ID"
// @Param        car  body      domain.Car    true  "Updated car data"
// @Success      200  {object}  domain.Car
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/cars/{id} [put]
func (h *HandlerImpl) UpdateCar(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid car ID"})
		return
	}

	var req domain.Car
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	car, _ := h.service.Update(id, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, car)
}

// DeleteCar godoc
// @Summary      Delete a car
// @Description  Deletes a car by ID
// @Tags         cars
// @Produce      json
// @Param        id   path      int  true  "Car ID"
// @Success      204  {string}  string  "No Content"
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /api/cars/{id} [delete]
func (h *HandlerImpl) DeleteCar(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid car ID"})
		return
	}

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "car deleted successfully"})
}
