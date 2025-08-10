package utils

import (
	"net/http"
	"user-management-golang/model"

	"github.com/gin-gonic/gin"
)

func JSONSuccess(c *gin.Context, data any, msg string) {
	c.JSON(http.StatusOK, model.ApiResponse{
		Status:  "success",
		Message: msg,
		Data:    data,
	})
}

func JSONError(c *gin.Context, code int, msg any) {
	c.JSON(code, model.ApiResponse{
		Status:  "error",
		Message: msg,
	})
}
