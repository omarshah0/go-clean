package middleware

import (
	"bytes"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type BodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

type LogEntry struct {
	Method       string
	IPAddress    string
	Trace        string
	Path         string
	UserType     string
	UserEmail    string
	Status       string
	StartTime    string
	EndTime      string
	Duration     string
	Payload      string
	ResponseBody string
}

func (w BodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func readBody(c *gin.Context) string {
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = io.ReadAll(c.Request.Body)
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	return string(bodyBytes)
}

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// Save a copy of this request for debugging.
		payload := readBody(c)

		// Write the response body to a buffer
		writer := &BodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = writer

		// Process request
		c.Next()

		// Log after processing the request
		endTime := time.Now()

		logEntry := &LogEntry{
			Method:       c.Request.Method,
			IPAddress:    c.ClientIP(),
			Path:         c.Request.URL.Path,
			Trace:        c.GetString("X-Trace-ID"),
			UserType:     c.GetString("user_type"),
			UserEmail:    c.GetString("user_email"),
			Status:       strconv.Itoa(c.Writer.Status()),
			StartTime:    startTime.Format(time.RFC3339),
			EndTime:      endTime.Format(time.RFC3339),
			Duration:     endTime.Sub(startTime).String(),
			Payload:      payload,
			ResponseBody: writer.body.String(),
		}

		go saveLogToFile(logEntry, "file.csv")
	}
}

func saveLogToFile(logEntry *LogEntry, filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{
		logEntry.Method,
		logEntry.IPAddress,
		logEntry.Trace,
		logEntry.Path,
		logEntry.UserType,
		logEntry.UserEmail,
		logEntry.Status,
		logEntry.StartTime,
		logEntry.EndTime,
		logEntry.Duration,
		logEntry.Payload,
		logEntry.ResponseBody,
	})
}
