package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	fmt.Errorf(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
