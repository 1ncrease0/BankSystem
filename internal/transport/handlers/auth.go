package handlers

import (
	"FinanceSystem/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) RegisterEnterpriseSpecialist(c *gin.Context) {
	enterpriseIDStr := c.PostForm("enterpriseId")
	username := c.PostForm("username")
	password := c.PostForm("password")

	if enterpriseIDStr == "" || username == "" || password == "" {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Заполните все поля.",
		})
		return
	}

	enterpriseID, err := strconv.Atoi(enterpriseIDStr)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Неверный формат ID предприятия.",
		})
		return
	}

	input := models.EnterpriseSpecialist{
		EnterpriseId: enterpriseID,
		UserName:     username,
		Password:     password,
	}

	_, err = h.services.Authorization.CreateEnterpriseSpecialist(input)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Ошибка регистрации. Попробуйте снова.",
		})
		return
	}

	if c.GetHeader("HX-Request") != "" {
		c.Header("HX-Redirect", "/v1/enterprise_specialist/login")
		c.Status(http.StatusOK)
		return
	}
	c.Redirect(http.StatusSeeOther, "/v1/enterprise_specialist/login")
}
func (h *Handler) LoginEnterpriseSpecialist(c *gin.Context) {
	enterpriseIDStr := c.PostForm("enterpriseId") // Изменено здесь
	username := c.PostForm("username")
	password := c.PostForm("password")

	if enterpriseIDStr == "" || username == "" || password == "" {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Заполните все поля.",
		})
		return
	}

	_, err := strconv.Atoi(enterpriseIDStr)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Неверный формат ID предприятия.",
		})
		return
	}

	token, err := h.services.Authorization.GenerateEnterpriseSpecialistToken(username, password)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Неверные учетные данные. Попробуйте снова.",
		})
		return
	}

	c.SetCookie("token", token, 3600, "/", "localhost", false, true)

	if c.GetHeader("HX-Request") != "" {
		c.Header("HX-Redirect", "/v1/enterprise_specialist/dashboard")
		c.Status(http.StatusOK)
		return
	}
	c.Redirect(http.StatusSeeOther, "/v1/enterprise_specialist/dashboard")
}

func (h *Handler) LoginEmployee(c *gin.Context) {

	bankID := c.Param("id")
	if bankID == "" {
		c.String(http.StatusBadRequest, "Bank ID is required")
		return
	}

	login := c.PostForm("login")
	password := c.PostForm("password")

	if login == "" || password == "" {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Неверные учетные данные. Попробуйте снова.",
		})
		return
	}

	token, err := h.services.Authorization.GenerateEmployeeToken(login, password)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Неверные учетные данные. Попробуйте снова.",
		})
		return
	}

	c.SetCookie("token", token, 3600, "/", "localhost", false, true)

	if c.GetHeader("HX-Request") != "" {
		c.Header("HX-Redirect", "/v1/employee/bank/"+bankID+"/dashboard")
		c.Status(http.StatusOK)
		return
	}
	c.Redirect(http.StatusSeeOther, "/v1/employee/bank/"+bankID+"/dashboard")
}

func (h *Handler) RegisterEmployee(c *gin.Context) {
	bankIDStr := c.Param("id")
	if bankIDStr == "" {
		c.String(http.StatusBadRequest, "Bank ID is required")
		return
	}

	bankID, err := strconv.Atoi(bankIDStr)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Неверный формат Bank ID.",
		})
		return
	}
	var input models.BankEmployee
	input.UserName = c.PostForm("user_name")
	input.Password = c.PostForm("password")
	input.Role = c.PostForm("role")
	input.BankId = bankID

	if input.UserName == "" || input.Password == "" {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Неверные учетные данные. Попробуйте снова.",
		})
		return
	}

	_, err = h.services.Authorization.CreateBankEmployee(input)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Неверные учетные данные. Попробуйте снова.",
		})
		return
	}

	if c.GetHeader("HX-Request") != "" {
		c.Header("HX-Redirect", "/v1/employee/bank/"+bankIDStr+"/login")
		c.Status(http.StatusOK)
		return
	}
	c.Redirect(http.StatusSeeOther, "/v1/employee/bank/"+bankIDStr+"/login")
}

func (h *Handler) LoginEmployeePage(c *gin.Context) {
	bankID := c.Param("id")

	if bankID == "" {
		c.String(http.StatusBadRequest, "Bank ID is required")
		return
	}

	c.HTML(http.StatusOK, "loginEmployeePage.html", gin.H{
		"BankID": bankID,
	})

}

func (h *Handler) RegisterEmployeePage(c *gin.Context) {
	bankID := c.Param("id")

	if bankID == "" {
		c.String(http.StatusBadRequest, "Bank ID is required")
		return
	}

	c.HTML(http.StatusOK, "registerEmployeePage.html", gin.H{
		"BankID": bankID,
	})
}

func (h *Handler) LoginClient(c *gin.Context) {
	login := c.PostForm("login")
	password := c.PostForm("password")

	if login == "" || password == "" {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Неверные учетные данные. Попробуйте снова.",
		})

		return
	}

	token, err := h.services.Authorization.GenerateClientToken(login, password)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Неверный логин или пароль. Попробуйте снова.",
		})
		return
	}
	c.SetCookie("token", token, 3600, "/", "localhost", false, true)

	if c.GetHeader("HX-Request") != "" {
		c.Header("HX-Redirect", "/v1/client/banks")
		c.Status(http.StatusOK)
		return
	}
	c.Redirect(http.StatusSeeOther, "/v1/client/banks")
}

func (h *Handler) RegisterClient(c *gin.Context) {

	var input models.Client

	input.Name = c.PostForm("name")
	input.Surname = c.PostForm("surname")
	input.Patronymic = c.PostForm("patronymic")
	input.UserName = c.PostForm("user_name")
	input.PassportSeries = c.PostForm("passport_series")
	input.PassportNumber = c.PostForm("passport_number")
	input.IdNumber = c.PostForm("id_number")
	input.PhoneNumber = c.PostForm("phone_number")
	input.Email = c.PostForm("email")
	input.Password = c.PostForm("password")

	if input.Name == "" || input.Surname == "" || input.UserName == "" || input.Password == "" {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Неверные учетные данные. Попробуйте снова.",
		})
		return
	}
	_, err := h.services.Authorization.CreateClient(input)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"error": "Неверные данные. Попробуйте снова.",
		})
		return
	}

	if c.GetHeader("HX-Request") != "" {
		c.Header("HX-Redirect", "/v1/client/login")
		c.Status(http.StatusOK)
		return
	}
	c.Redirect(http.StatusSeeOther, "/v1/client/login")
}

func (h *Handler) LoginEnterprisePage(c *gin.Context) {
	c.HTML(http.StatusOK, "loginEnterprisePage.html", nil)
}
func (h *Handler) RegisterEnterprisePage(c *gin.Context) {
	c.HTML(http.StatusOK, "registerEnterprisePage.html", nil)
}
func (h *Handler) RegisterClientPage(c *gin.Context) {
	c.HTML(http.StatusOK, "registerClientPage.html", nil)
}

func (h *Handler) LoginClientPage(c *gin.Context) {
	c.HTML(http.StatusOK, "loginClientPage.html", nil)
}

func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.Redirect(http.StatusSeeOther, "/v1/")
}

func (h *Handler) WelcomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "welcomePage.html", nil)
}
