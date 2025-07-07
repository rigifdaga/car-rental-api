package handlers

import (
    "car-rental-api/internal/models"
    "car-rental-api/internal/services"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type DriverIncentiveHandler struct {
    driverIncentiveService *services.DriverIncentiveService
}

func NewDriverIncentiveHandler(driverIncentiveService *services.DriverIncentiveService) *DriverIncentiveHandler {
    return &DriverIncentiveHandler{driverIncentiveService: driverIncentiveService}
}

func (h *DriverIncentiveHandler) CreateDriverIncentive(c *gin.Context) {
    var request struct {
        RentalID int `json:"rental_id" binding:"required"`
    }
    
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.driverIncentiveService.CreateDriverIncentive(request.RentalID); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Driver incentive created successfully",
    })
}

func (h *DriverIncentiveHandler) GetAllDriverIncentives(c *gin.Context) {
    incentives, err := h.driverIncentiveService.GetAllDriverIncentives()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Driver incentives retrieved successfully",
        "data":    incentives,
    })
}

func (h *DriverIncentiveHandler) GetDriverIncentiveByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid incentive ID"})
        return
    }

    incentive, err := h.driverIncentiveService.GetDriverIncentiveByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Driver incentive not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Driver incentive retrieved successfully",
        "data":    incentive,
    })
}

func (h *DriverIncentiveHandler) GetDriverIncentiveByRentalID(c *gin.Context) {
    rentalID, err := strconv.Atoi(c.Param("rental_id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rental ID"})
        return
    }

    incentive, err := h.driverIncentiveService.GetDriverIncentiveByRentalID(rentalID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Driver incentive not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Driver incentive retrieved successfully",
        "data":    incentive,
    })
}

func (h *DriverIncentiveHandler) UpdateDriverIncentive(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid incentive ID"})
        return
    }

    var incentive models.DriverIncentive
    if err := c.ShouldBindJSON(&incentive); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.driverIncentiveService.UpdateDriverIncentive(id, &incentive); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Driver incentive updated successfully",
    })
}

func (h *DriverIncentiveHandler) DeleteDriverIncentive(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid incentive ID"})
        return
    }

    if err := h.driverIncentiveService.DeleteDriverIncentive(id); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Driver incentive deleted successfully",
    })
}