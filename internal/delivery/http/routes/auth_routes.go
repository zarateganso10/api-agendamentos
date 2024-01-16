package routes

import (
	"api-agendamentos/internal/delivery/http/handlers"
	"api-agendamentos/internal/repository"
	"api-agendamentos/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func AuthRoutes(connectionPG *sqlx.DB, router *gin.Engine) {
	userRepository := repository.NewUserRepository(connectionPG)
	authService := services.NewAuthService(userRepository)
	authHandler := handlers.NewAuthHandler(authService)
	auth := router.Group("api/auth")
	{
		auth.POST("/", authHandler.CreateSession)
	}
}
