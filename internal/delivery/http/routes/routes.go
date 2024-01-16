package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func InitRoutes(connectionDB *sqlx.DB, router *gin.Engine) {
	UserRoutes(connectionDB, router)
	AuthRoutes(connectionDB, router)
	CompaniesRoutes(connectionDB, router)
	AppointmentsRoutes(connectionDB, router)
}
