package controller

import (
	"net/http"
	"user-management-golang/model"

	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Username string
	Password string
}
type userResponse struct {
	ID       uint
	Name     string
	Age      int
	Username string
}

type apiResponse struct {
	Status  string `json:"status"`
	Message any    `json:"message"`
	Data    any    `json:"data"`
}

func toResponse(user model.User) userResponse {
	var response userResponse
	response.ID = user.ID
	response.Name = user.Name
	response.Age = user.Age
	response.Username = user.Username
	return response
}

func jsonSuccess(c *gin.Context, data any, msg string) {
	c.JSON(http.StatusOK, apiResponse{
		Status:  "success",
		Message: msg,
		Data:    data,
	})
}

func jsonError(c *gin.Context, code int, msg any) {
	c.JSON(code, apiResponse{
		Status:  "error",
		Message: msg,
	})
}
