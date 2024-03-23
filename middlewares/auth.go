package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/trishan9/Gin-CRUD/helpers"
	"github.com/trishan9/Gin-CRUD/initializers"
	"github.com/trishan9/Gin-CRUD/models"
)

func IsAuthenticated(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	claims, err := helpers.VerifyToken(tokenString)

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	var user models.User
	initializers.DB.First(&user, claims["id"])

	if user.ID == 0 || !user.IsVerified {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	c.Set("user", user)
	c.Next()
}
