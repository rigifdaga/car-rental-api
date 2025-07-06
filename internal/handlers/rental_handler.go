package handlers

import (
    "car-rental-api/internal/models"
    "car-rental-api/internal/services"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type RentalHandler struct {
    rentalService *services.RentalService
}

func NewRentalHandler(rentalService *services.RentalService) *RentalHandler {
    return &RentalHandler{rentalService: rentalService}
}

func (h *RentalHandler) CreateRental(c *gin.Context) {
    var rental models.Rental
    if err := c.ShouldBindJSON(&rental); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.rentalService.CreateRental(&rental); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Rental created successfully",
        "data":    rental,
    })
}

func (h *RentalHandler) GetAllRentals(c *gin.Context) {
    rentals, err := h.rentalService.GetAllRentals()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Rentals retrieved successfully",
        "data":    rentals,
    })
}

func (h *RentalHandler) GetRentalByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rental ID"})
        return
    }

    rental, err := h.rentalService.GetRentalByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Rental not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Rental retrieved successfully",
        "data":    rental,
    })
}

func (h *RentalHandler) UpdateRental(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rental ID"})
        return
    }

    var rental models.Rental
    if err := c.ShouldBindJSON(&rental); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.rentalService.UpdateRental(id, &rental); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Rental updated successfully",
    })
}

func (h *RentalHandler) DeleteRental(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rental ID"})
        return
    }

    if err := h.rentalService.DeleteRental(id); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Rental deleted successfully",
    })
}

func (h *RentalHandler) GetActiveRentals(c *gin.Context) {
    rentals, err := h.rentalService.GetActiveRentals()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Active rentals retrieved successfully",
        "data":    rentals,
    })
}