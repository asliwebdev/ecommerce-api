package handler

import (
	"ecommerce/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userService    *service.UserService
	productService *service.ProductService
}

func NewHandler(userService *service.UserService, productService *service.ProductService) *Handler {
	return &Handler{
		userService:    userService,
		productService: productService,
	}
}

func Run(h *Handler) *gin.Engine {
	router := gin.Default()

	// USER ROUTES
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/", h.CreateUser)
		userRoutes.GET("/", h.GetAllUsers)
		userRoutes.GET("/:id", h.GetUserById)
		userRoutes.PUT("/:id", h.UpdateUser)
		userRoutes.DELETE("/:id", h.DeleteUser)
	}

	// PRODUCT ROUTES
	productRoutes := router.Group("/products")
	{
		productRoutes.POST("/", h.CreateProduct)
		productRoutes.GET("/", h.GetAllProducts)
		productRoutes.GET("/:id", h.GetProductById)
		productRoutes.PUT("/:id", h.UpdateProduct)
		productRoutes.DELETE("/:id", h.DeleteProduct)
	}

	return router
}
