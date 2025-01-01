package handler

import (
	"net/http"

	"ecommerce/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(ctx *gin.Context) {
	user := models.User{}
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.userService.CreateUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created"})
}
