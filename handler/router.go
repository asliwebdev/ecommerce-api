package handler

import (
	"ecommerce/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userService *service.UserService
}

func NewHandler(userService *service.UserService) *Handler {
	return &Handler{
		userService: userService,
	}
}

func Run(h *Handler) *gin.Engine {
	r := gin.Default()

	return r
}
