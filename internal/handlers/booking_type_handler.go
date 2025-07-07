package handlers

import (
    "car-rental-api/internal/models"
    "car-rental-api/internal/services"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type BookingTypeHandler struct {
    bookingTypeService *services.BookingTypeService
}

func NewBookingTypeHandler(bookingTypeService *services.BookingTypeService) *BookingTypeHandler {
    return &BookingTypeHandler{bookingTypeService: bookingTypeService}
}

func (h *BookingTypeHandler) CreateBookingType(c *gin.Context) {
    var bookingType models.BookingType
    if err := c.ShouldBindJSON(&bookingType); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.bookingTypeService.CreateBookingType(&bookingType); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Booking type created successfully",
        "data":    bookingType,
    })
}

func (h *BookingTypeHandler) GetAllBookingTypes(c *gin.Context) {
    bookingTypes, err := h.bookingTypeService.GetAllBookingTypes()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Booking types retrieved successfully",
        "data":    bookingTypes,
    })
}

func (h *BookingTypeHandler) GetBookingTypeByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking type ID"})
        return
    }

    bookingType, err := h.bookingTypeService.GetBookingTypeByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Booking type not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Booking type retrieved successfully",
        "data":    bookingType,
    })
}

func (h *BookingTypeHandler) UpdateBookingType(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking type ID"})
        return
    }

    var bookingType models.BookingType
    if err := c.ShouldBindJSON(&bookingType); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.bookingTypeService.UpdateBookingType(id, &bookingType); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Booking type updated successfully",
    })
}

func (h *BookingTypeHandler) DeleteBookingType(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid booking type ID"})
        return
    }

    if err := h.bookingTypeService.DeleteBookingType(id); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Booking type deleted successfully",
    })
}