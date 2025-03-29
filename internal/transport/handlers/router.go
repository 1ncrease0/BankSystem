package handlers

import (
	"FinanceSystem/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.LoadHTMLGlob("web/templates/*")

	v1 := r.Group("/v1")
	{
		v1.GET("/", h.WelcomePage)

		v1.GET("/logout", h.Logout)

		v1.GET("/client/register", h.RegisterClientPage)
		v1.POST("/client/register", h.RegisterClient)
		v1.GET("/client/login", h.LoginClientPage)
		v1.POST("/client/login", h.LoginClient)

		v1.GET("/banks", h.AllBanksPage)

		v1.GET("/employee/bank/:id/login", h.LoginEmployeePage)
		v1.GET("/employee/bank/:id/register", h.RegisterEmployeePage)
		v1.POST("/employee/bank/:id/login", h.LoginEmployee)
		v1.POST("/employee/bank/:id/register", h.RegisterEmployee)

		v1.GET("/enterprise_specialist/login", h.LoginEnterprisePage)
		v1.GET("/enterprise_specialist/register", h.RegisterEnterprisePage)

		v1.POST("/enterprise_specialist/login", h.LoginEnterpriseSpecialist)
		v1.POST("/enterprise_specialist/register", h.RegisterEnterpriseSpecialist)

		enterprise := v1.Group("/enterprise_specialist", h.enterpriseSpecialistIdentity)
		{
			enterprise.GET("/dashboard", h.enterpriseDashboardPage)
			enterprise.POST("/account/:id/transfer", h.enterpriseTransfer)
			enterprise.POST("/salary_project/create", h.SalaryProjectCreate)

		}

		employee := v1.Group("/employee", h.employeeIdentity)
		{
			employee.GET("/bank/:id/dashboard", h.employeeDashboardPage)
			employee.POST("/bank/:id/request/:reqId/approve", h.approveRegRequest)
			employee.POST("/bank/:id/request/:reqId/reject", h.rejectRegRequest)

			employee.POST("/bank/:id/salary-request/:requestId/approve", h.ApproveSalaryProjectRequest)
			employee.POST("/bank/:id/salary-request/:requestId/reject", h.RejectSalaryProjectRequest)
			employee.POST("/bank/:id/action-log/:logId/cancel", h.CancelTransfer)
		}

		client := v1.Group("/client", h.clientIdentity)
		{
			client.GET("/healthcheck", h.HealthCheck)
			client.GET("/banks", h.ClientBanksPage)
			client.GET("/bank/:id", h.ClientDashboardPage)
			client.POST("/bank/:id/register", h.RegistrationRequest)

			client.POST("/bank/:id/request/create", h.CreateRequest)
			client.POST("/bank/:id/account/create", h.CreateAccount)
			client.POST("/account/:id/plus", h.PlusMoney)
			client.POST("/account/:id/minus", h.MinusMoney)
			client.POST("/account/:id/transfer", h.Transfer)
		}

	}
	return r
}
