package routes

import (
	"api-agendamentos/internal/delivery/http/handlers"
	"api-agendamentos/internal/repository"
	"api-agendamentos/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func UserRoutes(connectionPG *sqlx.DB, router *gin.Engine) {
	userRepository := repository.NewUserRepository(connectionPG)
	companyRepository := repository.NewCompanyRepository(connectionPG)
	userService := services.NewUserService(userRepository, companyRepository)
	userHandler := handlers.NewUserHandler(userService)
	user := router.Group("api/users")
	{
		user.GET("/", userHandler.ListUsers)
		user.GET("/:id", userHandler.GetUserByID)
		user.POST("/", userHandler.CreateUser)
		// user.PUT("/:id", userHandler.UpdateUserHandler)
		user.DELETE("/:id", userHandler.DeleteUserByID)
	}
}
