package handlers

import (
    "car-rental-api/internal/models"
    "car-rental-api/internal/services"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type CarHandler struct {
    carService *services.CarService
}

func NewCarHandler(carService *services.CarService) *CarHandler {
    return &CarHandler{carService: carService}
}

func (h *CarHandler) CreateCar(c *gin.Context) {
    var car models.Car
    if err := c.ShouldBindJSON(&car); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.carService.CreateCar(&car); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Car created successfully",
        "data":    car,
    })
}

func (h *CarHandler) GetAllCars(c *gin.Context) {
    cars, err := h.carService.GetAllCars()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Cars retrieved successfully",
        "data":    cars,
    })
}

func (h *CarHandler) GetCarByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid car ID"})
        return
    }

    car, err := h.carService.GetCarByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Car retrieved successfully",
        "data":    car,
    })
}

func (h *CarHandler) UpdateCar(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid car ID"})
        return
    }

    var car models.Car
    if err := c.ShouldBindJSON(&car); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.carService.UpdateCar(id, &car); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Car updated successfully",
    })
}

func (h *CarHandler) DeleteCar(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid car ID"})
        return
    }

    if err := h.carService.DeleteCar(id); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Car deleted successfully",
    })
}

func (h *CarHandler) GetAvailableCars(c *gin.Context) {
    cars, err := h.carService.GetAvailableCars()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Available cars retrieved successfully",
        "data":    cars,
    })
}