package handlers

import (
	"api-agendamentos/internal/delivery/http/utils"
	"api-agendamentos/internal/dto"
	"api-agendamentos/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppointmentHandler struct {
	AppointmentService *services.AppointmentService
}

func NewAppointmentHandler(appointmentService *services.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{
		AppointmentService: appointmentService,
	}
}

func (handler *AppointmentHandler) ListAppointements(c *gin.Context) {
	queryParams := utils.QueryParamsFromListUsers(c)
	result, err := handler.AppointmentService.ListAppointements(queryParams.Page, queryParams.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

func (handler *AppointmentHandler) GetAppointmentByID(c *gin.Context) {
	id := c.Param("id")
	appointment, err := handler.AppointmentService.GetAppointmentByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, appointment)
}

func (handler *AppointmentHandler) CreateAppointment(c *gin.Context) {
	var newAppointment dto.CreateAppointmentInput
	if err := c.BindJSON(&newAppointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	appointmentCreated, errorResponse := handler.AppointmentService.CreateAppointment(newAppointment)
	if errorResponse != nil {
		c.JSON(errorResponse.Status, gin.H{"error": errorResponse.Message})
		return
	}
	c.IndentedJSON(http.StatusCreated, appointmentCreated)
}
