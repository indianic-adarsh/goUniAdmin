package middleware

import (
	"net/http"
	"strings"

	"goUniAdmin/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware verifies JWT tokens
func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "Authorization header required"})
			c.Abort()
			return
		}

		// Expect "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "Invalid Authorization header format"})
			c.Abort()
			return
		}

		tokenStr := parts[1]
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(cfg.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Extract admin ID from token claims and set it in context
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			adminID, ok := claims["id"].(string)
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "Invalid token claims"})
				c.Abort()
				return
			}
			c.Set("adminID", adminID)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "Invalid token claims"})
			c.Abort()
			return
		}

		c.Next()
	}
}
