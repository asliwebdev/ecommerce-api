package handler

import (
	"ecommerce/models"
	"ecommerce/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Checkout(c *gin.Context) {
	var request models.CheckoutRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if request.Status == "" {
		request.Status = "pending"
	}

	resp, err := h.checkoutService.Checkout(request)
	if err != nil {
		switch err {
		case service.ErrCartIsEmpty:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Your cart is empty"})
		case service.ErrFailedToFetchCart:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch cart items"})
		case service.ErrFailedToFetchProduct:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch product details"})
		case service.ErrFailedToCreateOrder:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create order"})
		case service.ErrFailedToClearCart:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not clear cart"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An unexpected error occurred"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order placed successfully",
		"order":   resp,
	})
}
