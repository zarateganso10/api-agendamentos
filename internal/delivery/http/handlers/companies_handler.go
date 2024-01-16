package handlers

import (
	"api-agendamentos/internal/delivery/http/utils"
	"api-agendamentos/internal/dto"
	"api-agendamentos/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CompanyHandler struct {
	CompanyService *services.CompanyService
}

func NewCompanyHandler(companyService *services.CompanyService) *CompanyHandler {
	return &CompanyHandler{
		CompanyService: companyService,
	}
}

func (handler *CompanyHandler) ListCompanies(c *gin.Context) {
	queryParams := utils.QueryParamsFromListUsers(c)
	result, err := handler.CompanyService.ListCompanies(queryParams.Page, queryParams.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

func (handler *CompanyHandler) GetCompanyByID(c *gin.Context) {
	id := c.Param("id")
	company, err := handler.CompanyService.GetCompanyByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, company)
}

func (handler *CompanyHandler) CreateCompany(c *gin.Context) {
	var newCompany dto.CreateCompanyInput
	if err := c.BindJSON(&newCompany); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	companyCreated, errorResponse := handler.CompanyService.CreateCompany(newCompany)
	if errorResponse != nil {
		c.JSON(errorResponse.Status, gin.H{"error": errorResponse.Message})
		return
	}
	c.IndentedJSON(http.StatusCreated, companyCreated)
}
