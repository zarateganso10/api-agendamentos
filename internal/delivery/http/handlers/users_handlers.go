package handlers

import (
	"api-agendamentos/internal/delivery/http/utils"
	"api-agendamentos/internal/dto"
	"api-agendamentos/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (handler *UserHandler) ListUsers(c *gin.Context) {
	queryParams := utils.QueryParamsFromListUsers(c)
	result, err := handler.UserService.ListUsers(queryParams.Page, queryParams.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, result)
}

func (handler *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := handler.UserService.GetUserByID(id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func (handler *UserHandler) CreateUser(c *gin.Context) {
	var newUser dto.CreateUserInput
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	userCreated, errorResponse := handler.UserService.CreateUser(newUser)
	if errorResponse != nil {
		c.JSON(errorResponse.Status, gin.H{"error": errorResponse.Message})
		return
	}
	c.IndentedJSON(http.StatusCreated, userCreated)
}

func (handler *UserHandler) DeleteUserByID(c *gin.Context) {
	id := c.Param("id")
	err := handler.UserService.DeleteUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.Writer.WriteHeader(http.StatusNoContent)
}
