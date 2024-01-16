package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type ListUsersParams struct {
	Page  int
	Limit int
}

func QueryParamsFromListUsers(c *gin.Context) *ListUsersParams {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "5"))

	return &ListUsersParams{
		Page:  page,
		Limit: limit,
	}
}
