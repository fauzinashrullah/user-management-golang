package middleware

import (
	"log"
	"strings"
	"user-management-golang/security"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.JSON(401, gin.H{"error": "Token missing"})
			c.Abort()
			return
		}

		var tokenString string
		if strings.HasPrefix(header, "Bearer ") {
			tokenString = strings.TrimPrefix(header, "Bearer ")
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return security.JwtSecret, nil
		})

		log.Println(token)
		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(401, gin.H{"error": "Invalid claims"})
			c.Abort()
			return
		}

		userID := claims["sub"].(string)
		c.Set("userID", userID)

		c.Next()
	}
}
