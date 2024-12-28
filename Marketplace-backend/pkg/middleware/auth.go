package middleware

import (
	"Marketplace-backend/internal/repository"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(tokenRepo repository.TokenRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		/// Get the token from authorization header
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			log.Println("Missing Authorization header")
			c.JSON(401, gin.H{"error": " Authorization token required"})
			c.Abort()
			return
		}
		// The token is usualy in the format  "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			log.Println("Invalid Authorization Format ")
			c.JSON(401, gin.H{"error": "Authorization format must be Bearer <token>"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Fetch the token from the database
		token, err := tokenRepo.FindByToken(tokenString)
		if err != nil {
			// log the token  lookup failure
			log.Printf("Token lookup failed: %v", err)
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Check if the token expired
		if token.ExpiredAt.Before(time.Now()) {
			log.Printf("Token expired at: %v", token.ExpiredAt)
			c.JSON(401, gin.H{"error": "Token expired"})
			c.Abort()
			return
		}
		// Set the user ID in the context
		c.Set("userID", token.UserID)

		// Proceed to the next hundler in the chain
		c.Next()

	}

}
