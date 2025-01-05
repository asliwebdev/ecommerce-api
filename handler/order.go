package handler

import (
	"ecommerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllOrders(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User Id is required"})
		return
	}

	orders, err := h.orderService.GetAllOrders(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch orders"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func (h *Handler) GetOrderById(c *gin.Context) {
	orderId := c.Param("orderId")

	order, err := h.orderService.GetOrderById(orderId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

func (h *Handler) UpdateOrder(c *gin.Context) {
	var req models.UpdateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	err := h.orderService.UpdateOrder(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully"})
}

func (h *Handler) DeleteOrder(c *gin.Context) {
	orderId := c.Param("orderId")

	err := h.orderService.DeleteOrder(orderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
