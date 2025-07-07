package handlers

import (
    "car-rental-api/internal/models"
    "car-rental-api/internal/services"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type DriverHandler struct {
    driverService *services.DriverService
}

func NewDriverHandler(driverService *services.DriverService) *DriverHandler {
    return &DriverHandler{driverService: driverService}
}

func (h *DriverHandler) CreateDriver(c *gin.Context) {
    var driver models.Driver
    if err := c.ShouldBindJSON(&driver); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.driverService.CreateDriver(&driver); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Driver created successfully",
        "data":    driver,
    })
}

func (h *DriverHandler) GetAllDrivers(c *gin.Context) {
    drivers, err := h.driverService.GetAllDrivers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Drivers retrieved successfully",
        "data":    drivers,
    })
}

func (h *DriverHandler) GetDriverByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid driver ID"})
        return
    }

    driver, err := h.driverService.GetDriverByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Driver not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Driver retrieved successfully",
        "data":    driver,
    })
}

func (h *DriverHandler) UpdateDriver(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid driver ID"})
        return
    }

    var driver models.Driver
    if err := c.ShouldBindJSON(&driver); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.driverService.UpdateDriver(id, &driver); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Driver updated successfully",
    })
}

func (h *DriverHandler) DeleteDriver(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid driver ID"})
        return
    }

    if err := h.driverService.DeleteDriver(id); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Driver deleted successfully",
    })
}