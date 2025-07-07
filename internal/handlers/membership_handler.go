package handlers

import (
    "car-rental-api/internal/models"
    "car-rental-api/internal/services"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
)

type MembershipHandler struct {
    membershipService *services.MembershipService
}

func NewMembershipHandler(membershipService *services.MembershipService) *MembershipHandler {
    return &MembershipHandler{membershipService: membershipService}
}

func (h *MembershipHandler) CreateMembership(c *gin.Context) {
    var membership models.Membership
    if err := c.ShouldBindJSON(&membership); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.membershipService.CreateMembership(&membership); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "message": "Membership created successfully",
        "data":    membership,
    })
}

func (h *MembershipHandler) GetAllMemberships(c *gin.Context) {
    memberships, err := h.membershipService.GetAllMemberships()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Memberships retrieved successfully",
        "data":    memberships,
    })
}

func (h *MembershipHandler) GetMembershipByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid membership ID"})
        return
    }

    membership, err := h.membershipService.GetMembershipByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Membership not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Membership retrieved successfully",
        "data":    membership,
    })
}

func (h *MembershipHandler) UpdateMembership(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid membership ID"})
        return
    }

    var membership models.Membership
    if err := c.ShouldBindJSON(&membership); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.membershipService.UpdateMembership(id, &membership); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Membership updated successfully",
    })
}

func (h *MembershipHandler) DeleteMembership(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid membership ID"})
        return
    }

    if err := h.membershipService.DeleteMembership(id); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Membership deleted successfully",
    })
}