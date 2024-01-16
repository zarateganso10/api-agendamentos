package handlers

import (
	"api-agendamentos/internal/entity"
	"api-agendamentos/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		AuthService: authService,
	}
}

func (service *AuthHandler) CreateSession(c *gin.Context) {
	var auth entity.Auth
	if err := c.BindJSON(&auth); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error to format json"})
		return
	}
	token, err := service.AuthService.CreateSession(&auth)
	if err != nil {
		c.JSON(err.Status, gin.H{"error": err.Message})
		return
	}

	c.IndentedJSON(http.StatusCreated, token)
}
