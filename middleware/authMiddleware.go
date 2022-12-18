package middleware

import (
	"net/http"

	"github.com/Gymkhana-Forms/backend/helpers"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Token")
		if clientToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "No token"})
			c.Abort()
			return
		}

		claims, msg := helpers.ValidateToken(clientToken)
		if msg != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			c.Abort()
			return
		}

		c.Set("Email", claims.Email)
		c.Set("Name", claims.Name)
		c.Set("RollNo", claims.RollNo)
		c.Set("User_type", claims.User_type)
		c.Next()

	}
}
