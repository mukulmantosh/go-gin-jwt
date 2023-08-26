package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-jwt/helpers"
	"net/http"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError,
				gin.H{"error": fmt.Sprintf("No Authorization Header provided")})
			c.Abort()
			return
		}

		claims, err := helpers.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("first_name", claims.FirstName)
		c.Set("last_name", claims.LastName)
		c.Set("uid", claims.UID)
		c.Set("user_type", claims.UserType)
		c.Next()

	}
}
