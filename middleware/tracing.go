package middleware

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func Tracing() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate a new trace ID for each request
		traceID := stringWithCharset(8, charset)

		// Add the trace ID to the request context
		c.Set("X-Trace-ID", traceID)

		// Set the trace ID in the response headers
		c.Writer.Header().Set("X-Trace-ID", traceID)

		// Continue processing other middleware and request handlers
		c.Next()

	}
}
