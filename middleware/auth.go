package middleware

import (
	"net/http"
	"strings"
	"user-management-golang/config"
	"user-management-golang/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			utils.JSONError(c, http.StatusUnauthorized, "Token missing")
			c.Abort()
			return
		}

		tokenString, _ := strings.CutPrefix(header, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return config.JwtSecret, nil
		})

		if err != nil || !token.Valid {
			utils.JSONError(c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.JSONError(c, http.StatusUnauthorized, "Invalid claims")
			c.Abort()
			return
		}

		userID := claims["sub"].(string)
		role := claims["role"].(string)
		c.Set("userID", userID)
		c.Set("role", role)

		c.Next()
	}
}

func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exist := c.Get("role")
		if !exist {
			utils.JSONError(c, http.StatusForbidden, "Forbidden")
			c.Abort()
			return
		}
		if userRole != role {
			utils.JSONError(c, http.StatusForbidden, "Forbidden")
			c.Abort()
			return
		}
		c.Next()
	}
}
