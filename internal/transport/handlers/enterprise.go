package handlers

import (
	"FinanceSystem/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) SalaryProjectCreate(c *gin.Context) {
	
	clientAccountID := c.PostForm("client_account_id")
	enterpriseAccountID := c.PostForm("enterprise_account_id")
	amountStr := c.PostForm("amount")

	
	if clientAccountID == "" || enterpriseAccountID == "" || amountStr == "" {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": "Все поля обязательны"})
		return
	}

	
	clientID, err := strconv.Atoi(clientAccountID)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": "Неверный ID клиента"})
		return
	}

	enterpriseID, err := strconv.Atoi(enterpriseAccountID)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": "Неверный ID предприятия"})
		return
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil || amount <= 0 {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": "Неверная сумма"})
		return
	}

	req := models.SalaryProjectRequest{
		ClientAccountId:     clientID,
		EnterpriseAccountId: enterpriseID,
		Amount:              amount,
		Status:              models.RequestUnderConsideration,
	}


	if _, err := h.services.SalaryProjectRequest.CreateSalaryProjectRequest(req); err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": "Ошибка создания: " + err.Error()})
		return
	}

	
	c.Redirect(http.StatusFound, "/v1/enterprise_specialist/dashboard")
}
func (h *Handler) enterpriseTransfer(c *gin.Context) {
	idS := c.Param("id")
	id, err := strconv.Atoi(idS)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный формат id")
		return
	}

	amountS := c.PostForm("amount")
	amount, err := strconv.ParseFloat(amountS, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный формат денег")
		return
	}

	recipientS := c.PostForm("recipientAccount")
	resepient, err := strconv.Atoi(recipientS)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный формат счета")
		return
	}
	err = h.services.Account.TransferMoney(id, resepient, amount)
	if err != nil {
		c.String(http.StatusBadRequest, "Ошибка")
		return
	}
	c.Redirect(http.StatusSeeOther, "/v1/enterprise_specialist/dashboard")
}
func (h *Handler) enterpriseDashboardPage(c *gin.Context) {

	enterpriseIDInterface, exists := c.Get(enterpriseSpecialistCtx)
	if !exists {
		c.Redirect(http.StatusSeeOther, "/v1/")
		return
	}

	enterpriseID, ok := enterpriseIDInterface.(int)
	if !ok {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "Неверные данные пользователя.",
		})
		return
	}
	enterpriseSpec, err := h.services.Authorization.EnterpriseSpecialist(enterpriseID)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{})
		return
	}

	accounts, err := h.services.Account.AccountsByEnterprise(enterpriseSpec.EnterpriseId)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "Ошибка получения счетов предприятия.",
		})
		return
	}

	salaryProjects, err := h.services.SalaryProject.SalaryProjectByEnterprise(enterpriseSpec.EnterpriseId)
	c.HTML(http.StatusOK, "dashboardEnterprisePage.html", gin.H{
		"accounts":       accounts,
		"salaryProjects": salaryProjects,
	})
}
