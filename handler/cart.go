package handler

import (
	"ecommerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddProduct(c *gin.Context) {
	var cart models.AddProductRequest

	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.cartService.AddProduct(&cart); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product successfully added to cart", "cartItem": cart})
}

func (h *Handler) GetCart(c *gin.Context) {
	userId := c.Param("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User Id is required"})
		return
	}

	cart, err := h.cartService.GetCart(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cart": cart})
}

func (h *Handler) RemoveProductFromCart(c *gin.Context) {
	userId := c.Param("userId")
	productId := c.Param("productId")

	if userId == "" || productId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id and product_id are required"})
		return
	}

	if err := h.cartService.RemoveProduct(userId, productId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product removed from cart"})
}

func (h *Handler) ClearCart(c *gin.Context) {
	userId := c.Param("userId")

	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	if err := h.cartService.ClearCart(userId); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart cleared"})
}

func (h *Handler) UpdateProductQuantity(c *gin.Context) {
	var cart models.UpdateProductQuantity

	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.cartService.UpdateProductQuantity(&cart); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product quantity successfully updated"})
}
