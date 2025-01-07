package handler

import (
	"ecommerce/middleware"
	"ecommerce/repository"
	"ecommerce/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userService     *service.UserService
	productService  *service.ProductService
	cartService     *service.CartService
	checkoutService *service.CheckoutService
	orderService    *service.OrderService
}

func NewHandler(userService *service.UserService, productService *service.ProductService, cartService *service.CartService, checkoutService *service.CheckoutService, orderService *service.OrderService) *Handler {
	return &Handler{
		userService:     userService,
		productService:  productService,
		cartService:     cartService,
		checkoutService: checkoutService,
		orderService:    orderService,
	}
}

func Run(h *Handler, userRepo *repository.UserRepo) *gin.Engine {
	router := gin.Default()

	router.POST("/users", h.CreateUser)

	router.Use(middleware.BasicAuth(userRepo))

	// USER ROUTES
	userRoutes := router.Group("/users")
	{
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

	// CART ROUTES
	cartRoutes := router.Group("/carts")
	{
		cartRoutes.POST("/", h.AddProduct)
		cartRoutes.GET("/:userId", h.GetCart)
		cartRoutes.DELETE("/remove/:userId/:productId", h.RemoveProductFromCart)
		cartRoutes.DELETE("/clear/:userId", h.ClearCart)
		cartRoutes.PUT("/", h.UpdateProductQuantity)
	}

	// CHECKOUT ROUTES
	router.POST("/checkout", h.Checkout)

	// ORDER ROUTES
	orderRoutes := router.Group("/orders")
	{
		orderRoutes.GET("/user/:userId", h.GetAllOrders)
		orderRoutes.GET("/:orderId", h.GetOrderById)
		orderRoutes.PUT("/", h.UpdateOrder)
		orderRoutes.DELETE("/:orderId", h.DeleteOrder)
	}

	return router
}
