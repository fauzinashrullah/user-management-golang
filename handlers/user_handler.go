package handlers

import (
	"net/http"
	"user-management-golang/services"
	"user-management-golang/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(s services.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) CurrentUser(c *gin.Context) {
	userID, exist := c.Get("userID")
	if !exist {
		utils.JSONError(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	response, err := h.service.GetUserDetail(userID.(string))
	if err != nil {
		utils.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONSuccess(c, response, "Get user detail success")
}

func (h *UserHandler) GetUser(c *gin.Context) {
	name := c.Query("name")

	response := h.service.GetUser(name)
	utils.JSONSuccess(c, response, "Get user success")
}

func (h *UserHandler) GetUserDetail(c *gin.Context) {
	id := c.Param("id")
	response, err := h.service.GetUserDetail(id)
	if err != nil {
		utils.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONSuccess(c, response, "Get user detail success")
}
