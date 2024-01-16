package routes

import (
	"api-agendamentos/internal/delivery/http/handlers"
	"api-agendamentos/internal/repository"
	"api-agendamentos/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func AppointmentsRoutes(connectionPG *sqlx.DB, router *gin.Engine) {
	appointmentRepository := repository.NewAppointmentRepository(connectionPG)
	appointmentService := services.NewAppointmentService(appointmentRepository)
	appointmentHandler := handlers.NewAppointmentHandler(appointmentService)
	appointment := router.Group("api/appointments")
	{
		appointment.GET("/", appointmentHandler.ListAppointements)
		appointment.GET("/:id", appointmentHandler.GetAppointmentByID)
		appointment.POST("/", appointmentHandler.CreateAppointment)
		// company.PUT("/:id", companyHandler.UpdatecompanyHandler)
	}
}
