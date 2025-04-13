package handlers

import (
	"FinanceSystem/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CancelTransfer(c *gin.Context) {

	logIDStr := c.Param("logId")
	logID, err := strconv.Atoi(logIDStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Некорректный ID лога")
		return
	}

	logRecord, err := h.services.ActionLog.GetByID(logID)
	if err != nil {
		c.String(http.StatusNotFound, "Лог перевода не найден")
		return
	}

	err = h.services.Account.CancelTransfer(logID, *logRecord.Sender, *logRecord.Recipient, *logRecord.Amount)
	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка при отмене перевода")
		return
	}

	c.String(http.StatusOK, "")
}
func (h *Handler) employeeDashboardPage(c *gin.Context) {
	bankIdStr := c.Param("id")
	bankId, err := strconv.Atoi(bankIdStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{"error": "Неверный ID банка"})
		return
	}

	requests, err := h.services.RegistrationRequest.RegistrationRequestsByBank(bankId)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": "Ошибка загрузки заявок на регистрацию"})
		return
	}

	finRequests, err := h.services.FinRequest.GetFinRequestsByBank(bankId)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": "Ошибка загрузки финансовых заявок"})
		return
	}

	salaryRequests, err := h.services.SalaryProjectRequest.GetSalaryProjectRequestsByBank(bankId)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": "Ошибка загрузки запросов на зарплатные проекты"})
		return
	}

	logs, err := h.services.ActionLog.GetByBankID(bankId)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{"error": "Ошибка загрузки логов операций"})
		return
	}

	data := gin.H{
		"bankId":         bankId,
		"requests":       requests,
		"finRequests":    finRequests,
		"salaryRequests": salaryRequests,
		"logs":           logs, 
	}

	c.HTML(http.StatusOK, "dashboardEmployeePage.html", data)
}

func (h *Handler) ApproveSalaryProjectRequest(c *gin.Context) {
	bankIDStr := c.Param("id")
	requestIDStr := c.Param("requestId")

	bankID, err := strconv.Atoi(bankIDStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный ID банка")
		return
	}

	requestID, err := strconv.Atoi(requestIDStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный ID запроса")
		return
	}

	if err := h.services.SalaryProjectRequest.ApproveRequest(requestID, bankID); err != nil {
		c.String(http.StatusInternalServerError, "Ошибка при одобрении: "+err.Error())
		return
	}

	c.String(http.StatusOK, models.RequestApproved)
}

func (h *Handler) RejectSalaryProjectRequest(c *gin.Context) {
	bankIDStr := c.Param("bankId")
	requestIDStr := c.Param("requestId")

	bankID, err := strconv.Atoi(bankIDStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный ID банка")
		return
	}

	requestID, err := strconv.Atoi(requestIDStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный ID запроса")
		return
	}

	if err := h.services.SalaryProjectRequest.RejectRequest(requestID, bankID); err != nil {
		c.String(http.StatusInternalServerError, "Ошибка при отклонении: "+err.Error())
		return
	}

	c.String(http.StatusOK, models.RequestRejected)
}

func (h *Handler) approveRegRequest(c *gin.Context) {
	bankIdStr := c.Param("id")
	reqIdStr := c.Param("reqId")

	bankId, err := strconv.Atoi(bankIdStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный ID банка")
		return
	}
	reqId, err := strconv.Atoi(reqIdStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный ID заявки")
		return
	}

	err = h.services.RegistrationRequest.ApproveRequest(reqId, bankId)
	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка одобрения заявки")
		fmt.Println(err)
		return
	}

	c.String(http.StatusOK, models.RequestApproved)
}

func (h *Handler) rejectRegRequest(c *gin.Context) {
	bankIdStr := c.Param("id")
	reqIdStr := c.Param("reqId")

	_, err := strconv.Atoi(bankIdStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный ID банка")
		return
	}
	reqId, err := strconv.Atoi(reqIdStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный ID заявки")
		return
	}

	err = h.services.RegistrationRequest.RejectRequest(reqId, reqId)
	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка отклонения заявки")
		return
	}

	c.String(http.StatusOK, models.RequestRejected)
}
