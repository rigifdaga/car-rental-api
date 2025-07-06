package handlers

import (
    "car-rental-api/internal/models"
    "car-rental-api/internal/services"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type CustomerHandler struct {
    customerService *services.CustomerService
}

func NewCustomerHandler(customerService *services.CustomerService) *CustomerHandler {
    return &CustomerHandler{customerService: customerService}
}

func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
    var customer models.Customer
    if err := c.ShouldBindJSON(&customer); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.customerService.CreateCustomer(&customer); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Customer created successfully",
        "data":    customer,
    })
}

func (h *CustomerHandler) GetAllCustomers(c *gin.Context) {
    customers, err := h.customerService.GetAllCustomers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Customers retrieved successfully",
        "data":    customers,
    })
}

func (h *CustomerHandler) GetCustomerByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
        return
    }

    customer, err := h.customerService.GetCustomerByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Customer retrieved successfully",
        "data":    customer,
    })
}

func (h *CustomerHandler) UpdateCustomer(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
        return
    }

    var customer models.Customer
    if err := c.ShouldBindJSON(&customer); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.customerService.UpdateCustomer(id, &customer); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Customer updated successfully",
    })
}

func (h *CustomerHandler) DeleteCustomer(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
        return
    }

    if err := h.customerService.DeleteCustomer(id); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Customer deleted successfully",
    })
}