package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	tokenCookie             = "token"
	clientCtx               = "client_id"
	employeeCtx             = "employee_id"
	enterpriseSpecialistCtx = "enterprise_specialist_id"
	roleCtx                 = "user_role"
)

func (h *Handler) clientIdentity(c *gin.Context) {

	token, err := c.Cookie(tokenCookie)
	if err != nil {
		clearAuthCookie(c)
		c.Redirect(http.StatusSeeOther, "/v1/client/login")
		c.Abort()
		return
	}

	claims, err := h.services.Authorization.ParseToken(token)
	if err != nil {
		clearAuthCookie(c)
		c.Redirect(http.StatusSeeOther, "/v1/client/login")
		c.Abort()
		return
	}

	if claims.Role != "client" {
		c.Redirect(http.StatusSeeOther, "/v1/client/login")
		c.Abort()
		return
	}
	c.Set(clientCtx, claims.UserId)
	c.Set(roleCtx, claims.Role)
	c.Next()
}

func clearAuthCookie(c *gin.Context) {
	c.SetCookie(tokenCookie, "", -1, "/", "", false, true)
}

func (h *Handler) employeeIdentity(c *gin.Context) {
	token, err := c.Cookie(tokenCookie)
	if err != nil {
		clearAuthCookie(c)
		c.Redirect(http.StatusSeeOther, "/v1/login")
		c.Abort()
		return
	}

	claims, err := h.services.Authorization.ParseToken(token)
	if err != nil {
		clearAuthCookie(c)
		c.Redirect(http.StatusSeeOther, "/v1/login")
		c.Abort()
		return
	}

	if claims.Role != "administrator" && claims.Role != "manager" && claims.Role != "operator" {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		c.Abort()
		return
	}
	c.Set(roleCtx, claims.Role)
	c.Set(employeeCtx, claims.UserId)
	c.Next()
}

func (h *Handler) enterpriseSpecialistIdentity(c *gin.Context) {
	token, err := c.Cookie(tokenCookie)
	if err != nil {
		clearAuthCookie(c)
		c.Redirect(http.StatusSeeOther, "/v1/")
		c.Abort()
		return
	}

	claims, err := h.services.Authorization.ParseToken(token)
	if err != nil {
		clearAuthCookie(c)
		c.Redirect(http.StatusSeeOther, "/v1/")
		c.Abort()
		return
	}

	if claims.Role != "enterprise_specialist" {
		c.JSON(http.StatusForbidden, gin.H{"error": "access denied"})
		c.Abort()
		return
	}
	c.Set(enterpriseSpecialistCtx, claims.UserId)
	c.Set(roleCtx, claims.Role)
	c.Next()
}
