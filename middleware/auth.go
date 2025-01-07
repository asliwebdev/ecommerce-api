package middleware

import (
	"ecommerce/repository"
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func BasicAuth(userRepo *repository.UserRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Basic ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or invalid format"})
			c.Abort()
			return
		}

		payload, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(authHeader, "Basic "))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
			c.Abort()
			return
		}

		credentials := strings.SplitN(string(payload), ":", 2)
		if len(credentials) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials format"})
			c.Abort()
			return
		}

		email := credentials[0]
		password := credentials[1]

		user, err := userRepo.GetUserByEmail(email)
		if err != nil || user.Password != password {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			c.Abort()
			return
		}

		c.Next()
	}
}
