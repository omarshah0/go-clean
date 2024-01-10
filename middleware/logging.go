package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// Process request
		c.Next()

		// Log after processing the request
		endTime := time.Now()
		log.Printf("HTTP method: %s, path: %s, user_type: %s, email: %s, status: %d, start time: %s, end time: %s, duration: %s",
			c.Request.Method,
			c.Request.URL.Path,
			c.GetString("user_type"),
			c.GetString("user_email"),
			c.Writer.Status(),
			startTime.Format(time.RFC3339),
			endTime.Format(time.RFC3339),
			endTime.Sub(startTime),
		)
	}
}
