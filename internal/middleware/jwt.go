package middleware

import (
	"net/http"
	"strings"

	"order-tracking/pkg/response"
	"order-tracking/pkg/utils"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, response.NewErrorResponse("Authorization header is required"))
			c.Abort()
			return
		}

		// Pastikan formatnya "Bearer token"
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, response.NewErrorResponse("Authorization header format must be Bearer {token}"))
			c.Abort()
			return
		}

		tokenString := parts[1]

		claims, err := utils.ParseJWTToken(tokenString, secret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, response.NewErrorResponse("Invalid token"))
			c.Abort()
			return
		}

		userID, ok := claims["user_id"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, response.NewErrorResponse("Invalid token claims"))
			c.Abort()
			return
		}

		role, ok := claims["role"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, response.NewErrorResponse("Invalid token claims"))
			c.Abort()
			return
		}

		c.Set("userID", uint(userID))
		c.Set("role", role)
		c.Next()
	}
}
