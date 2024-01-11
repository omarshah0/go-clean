package middleware

import (
	"bytes"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// Save a copy of this request for debugging.
		payload := readBody(c)

		// Write the response body to a buffer
		writer := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = writer

		// Process request
		c.Next()

		// Log after processing the request
		endTime := time.Now()
		log.Printf("HTTP method: %s, path: %s, user_type: %s, email: %s, status: %d, start time: %s, end time: %s, duration: %s, payload: %s, response body: %s",
			c.Request.Method,
			c.Request.URL.Path,
			c.GetString("user_type"),
			c.GetString("user_email"),
			c.Writer.Status(),
			startTime.Format(time.RFC3339),
			endTime.Format(time.RFC3339),
			endTime.Sub(startTime),
			payload,
			writer.body.String(),
		)
	}
}

func readBody(c *gin.Context) string {
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = io.ReadAll(c.Request.Body)
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	return string(bodyBytes)
}
