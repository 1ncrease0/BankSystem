package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) HealthCheck(c *gin.Context) {
	c.HTML(http.StatusOK, "healthcheckPage.html", nil)
}
