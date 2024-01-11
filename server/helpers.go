package server

import (
	"github.com/gin-gonic/gin"

	"github.com/omarshah0/go-clean-architecture/types"
)

func getTraceID(c *gin.Context) string {
	traceID, exists := c.Get("X-Trace-ID")
	if !exists {
		return "unknown"
	}

	traceIDStr, ok := traceID.(string)
	if !ok {
		return "unknown"
	}

	return traceIDStr
}

func sendErrorResponse(c *gin.Context, err *types.HandlerErrorResponse) {
	traceID := getTraceID(c)
	c.JSON(err.StatusCode, &types.ErrorResponse{
		Type:    err.Type,
		Trace:   traceID,
		Message: err.Message,
		Error:   err.Error,
	})
}

func sendSuccessResponse(c *gin.Context, data interface{}, statusCode int) {
	c.JSON(statusCode, &types.SuccessResponse{
		Data: data,
	})
}

// Not Found Routes
func handleNotFoundRoute(c *gin.Context) {
	sendErrorResponse(c, &types.HandlerErrorResponse{
		Type:       "NotFound",
		Message:    "Not Found",
		Error:      "Endpoint not defined",
		StatusCode: 404,
	})
}
