package handlers

import (
	"FinanceSystem/internal/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) Transfer(c *gin.Context) {
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
	//c.Redirect(http.StatusSeeOther, "/v1/enterprise_specialist/dashboard")
}

func (h *Handler) PlusMoney(c *gin.Context) {
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
	err = h.services.Account.PutMoney(id, amount)
	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка сервера")
		return
	}
}
func (h *Handler) MinusMoney(c *gin.Context) {
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
	err = h.services.Account.WithdrawMoney(id, amount)
	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка сервера")
		return
	}
}
func (h *Handler) CreateAccount(c *gin.Context) {
	val, ok := c.Get(clientCtx)
	if !ok {
		c.HTML(http.StatusUnauthorized, "error.html", gin.H{
			"error": "Client id not found",
		})
		return
	}

	clientId, ok := val.(int)
	if !ok {
		c.HTML(http.StatusUnauthorized, "error.html", gin.H{
			"error": "Client id not found",
		})
		return
	}
	bankIdS := c.Param("id")
	bankId, err := strconv.Atoi(bankIdS)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный формат id банка")
		return
	}

	account := models.Account{
		ClientId:     &clientId,
		EnterpriseId: nil,
		BankId:       bankId,
		Currency:     "RUB",
		Balance:      0,
		Status:       models.StatusAvailable,
		LastUpdate:   time.Now(),
	}
	_, err = h.services.Account.CreateAccount(account)
	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка сервера")
		return
	}
}
func (h *Handler) BlockAccount(c *gin.Context) {
	bankIdS := c.Param("id")
	_, err := strconv.Atoi(bankIdS)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный формат id банка")
		return
	}
	accountIdS := c.Param("accountId")
	accountId, err := strconv.Atoi(accountIdS)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный формат id счета")
		return
	}
	err = h.services.Account.ChangeStatus(accountId, models.StatusBlocked)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный формат id счета")
		return
	}
}

func (h *Handler) FreezeAccount(c *gin.Context) {
	bankIdS := c.Param("id")
	_, err := strconv.Atoi(bankIdS)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный формат id банка")
		return
	}
	accountIdS := c.Param("accountId")
	accountId, err := strconv.Atoi(accountIdS)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный формат id счета")
		return
	}
	err = h.services.Account.ChangeStatus(accountId, models.StatusFrozen)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный формат id счета")
		return
	}
}

func (h *Handler) CreateRequest(c *gin.Context) {
	val, ok := c.Get(clientCtx)
	if !ok {
		c.HTML(http.StatusUnauthorized, "error.html", gin.H{
			"error": "Client id not found",
		})
		return
	}

	clientId, ok := val.(int)
	if !ok {
		c.HTML(http.StatusUnauthorized, "error.html", gin.H{
			"error": "Client id not found",
		})
		return
	}

	bankIdS := c.Param("id")
	bankId, err := strconv.Atoi(bankIdS)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный формат id банка")
		return
	}

	requestType := c.PostForm("requestType")
	accountIdStr := c.PostForm("accountId")
	amountStr := c.PostForm("amount")
	interestRateStr := c.PostForm("interestRate")
	termMonthsStr := c.PostForm("termMonths")

	if requestType == "" || accountIdStr == "" || amountStr == "" ||
		interestRateStr == "" || termMonthsStr == "" {
		c.String(http.StatusBadRequest, "Все поля обязательны для заполнения")
		return
	}

	accountId, err := strconv.Atoi(accountIdStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный формат accountId")
		return
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный формат суммы")
		return
	}

	interestRate, err := strconv.ParseFloat(interestRateStr, 64)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный формат процентной ставки")
		return
	}

	termMonths, err := strconv.Atoi(termMonthsStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный формат срока")
		return
	}

	finR := models.FinRequest{
		Status:       models.RequestUnderConsideration,
		Type:         requestType,
		ClientId:     clientId,
		BankId:       bankId,
		AccountId:    accountId,
		Amount:       amount,
		InterestRate: interestRate,
		TermMonths:   termMonths,
		CreatedAt:    time.Now(),
	}
	fmt.Println(finR)
	_, err = h.services.FinRequest.CreateFinRequest(finR)
	if err != nil {
		fmt.Println(err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// Check if the request was made with HTMX
	if c.GetHeader("HX-Request") == "true" {
		// Return a success message without replacing the entire modal
		c.HTML(http.StatusOK, "request-success.html", gin.H{
			"message": "Заявка успешно создана",
		})
	} else {
		// For non-HTMX requests, redirect back to the dashboard
		c.Redirect(http.StatusSeeOther, fmt.Sprintf("/v1/client/bank/%d", bankId))
	}
}

func (h *Handler) RegistrationRequest(c *gin.Context) {
	bankIdStr := c.Param("id")
	if bankIdStr == "" {
		c.String(http.StatusBadRequest, "Отсутствует id банка")
		return
	}
	bankId, err := strconv.Atoi(bankIdStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Неверный формат id банка")
		return
	}

	val, ok := c.Get(clientCtx)
	if !ok {
		c.String(http.StatusUnauthorized, "Client id not found")
		return
	}
	clientId, ok := val.(int)
	if !ok {
		c.String(http.StatusInternalServerError, "Некорректный client id")
		return
	}

	accounts, err := h.services.Account.AccountsByClientAndBank(clientId, bankId)
	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка при проверке счетов: "+err.Error())
		return
	}
	if len(accounts) > 0 {
		c.Redirect(http.StatusSeeOther, fmt.Sprintf("/v1/client/bank/%d", bankId))
		return
	}

	req := models.RegistrationRequest{
		ClientId:  clientId,
		BankId:    bankId,
		Status:    models.RequestUnderConsideration,
		CreatedAt: time.Now(),
	}
	_, err = h.services.RegistrationRequest.CreateRegistrationRequest(req)
	if err != nil {
		c.String(http.StatusInternalServerError, "Ошибка при создании заявки: "+err.Error())
		return
	}

	c.String(http.StatusOK, "Запрос на регистрацию успешно создан. Пожалуйста, ожидайте одобрения банка.")
}

func (h *Handler) AllBanksPage(c *gin.Context) {

	banks, err := h.services.Bank.Banks()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "Internal error",
		})
		return
	}

	c.HTML(http.StatusOK, "allBanksPage.html", gin.H{
		"Banks": banks,
	})

}

func (h *Handler) ClientBanksPage(c *gin.Context) {
	val, ok := c.Get(clientCtx)
	if !ok {
		c.HTML(http.StatusUnauthorized, "error.html", gin.H{
			"error": "Client id not found",
		})
		return
	}

	clientId, ok := val.(int)
	if !ok {
		c.HTML(http.StatusUnauthorized, "error.html", gin.H{
			"error": "Client id not found",
		})
		return
	}
	banks, err := h.services.Bank.FilteredByClient(clientId)

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "Internal error",
		})
		return
	}

	c.HTML(http.StatusOK, "clientBanksPage.html", gin.H{
		"filteredBanks": banks,
	})

}
func (h *Handler) ClientDashboardPage(c *gin.Context) {
	bankIdStr, ok := c.Params.Get("id")
	if !ok {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": "Bad request: отсутствует id банка",
		})
		return
	}
	bankId, err := strconv.Atoi(bankIdStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": "Bad request: неверный формат id банка",
		})
		return
	}

	val, ok := c.Get(clientCtx)
	if !ok {
		c.HTML(http.StatusUnauthorized, "error.html", gin.H{
			"error": "Client id not found",
		})
		return
	}
	clientId, ok := val.(int)
	if !ok {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "Некорректный тип client id",
		})
		return
	}

	info, err := h.services.ClientFinance.ClientFinanceInfoByBank(clientId, bankId)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	client, err := h.services.Authorization.Client(clientId)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "Internal error  h.services.Authorization.Client(clientId)",
		})
		return
	}

	bank, err := h.services.Bank.Bank(bankId)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": "Internal error h.services.Bank.Bank(bankId)",
		})
		return
	}

	data := models.DashboardData{
		Accounts:     info.Accounts,
		Credits:      info.Credits,
		Deposits:     info.Deposits,
		Installments: info.Installments,
		Client:       client,
		Bank:         bank,
	}

	c.HTML(http.StatusOK, "dashboardClientPage.html", data)
}
