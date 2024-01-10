package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/omarshah0/go-clean-architecture/types"
)

var secrets = map[string]string{
	"customer": "customer_secret",
	"driver":   "driver_secret",
	"admin":    "admin_secret",
}

func AuthMiddleware(userType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			return
		}

		// Check if the Authorization header starts with "Bearer "
		if len(authHeader) > 7 && strings.HasPrefix(authHeader, "Bearer ") {
			authHeader = authHeader[7:] // Remove "Bearer " prefix
		}

		// Verify the token using the secret key for the user type
		secret, ok := secrets[userType]
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Invalid user type"})
			return
		}

		// Parse the token
		token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
			// Replace "your-secret-key" with the actual secret key used to sign the JWT
			return []byte(secret), nil
		})

		// Check for parsing errors
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Check if the token is valid
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Retrieve user_id from the claims
			userID, ok := claims["user_id"]
			userEmail, ok := claims["user_email"]
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user_id or email in token"})
				c.Abort()
				return
			}

			// Add user_id to the request context
			c.Writer.Header().Set("X-User-Type", userType)
			c.Set("user_type", userType)
			c.Set("user_id", userID)
			c.Set("user_email", userEmail)

			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}
	}
}

// GenerateToken generates a JWT token with the given user type
func GenerateToken(user *types.User) (string, error) {
	// Get the secret key for the user type
	secret, ok := secrets[string(user.Type)]
	if !ok {
		return "", fmt.Errorf("Invalid user type")
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    1,
		"user_email": user.Email,
		"user_type":  user.Type,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign the token with the secret key
	return token.SignedString([]byte(secret))
}
