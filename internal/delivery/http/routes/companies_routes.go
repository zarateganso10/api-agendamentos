package routes

import (
	"api-agendamentos/internal/delivery/http/handlers"
	"api-agendamentos/internal/repository"
	"api-agendamentos/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func CompaniesRoutes(connectionPG *sqlx.DB, router *gin.Engine) {
	companyRepository := repository.NewCompanyRepository(connectionPG)
	companyService := services.NewCompanyService(companyRepository)
	companyHandler := handlers.NewCompanyHandler(companyService)
	company := router.Group("api/companies")
	{
		company.GET("/", companyHandler.ListCompanies)
		company.GET("/:id", companyHandler.GetCompanyByID)
		company.POST("/", companyHandler.CreateCompany)
		// company.PUT("/:id", companyHandler.UpdatecompanyHandler)
	}
}
