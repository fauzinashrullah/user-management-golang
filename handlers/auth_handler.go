package handlers

import (
	"net/http"
	"user-management-golang/model"
	"user-management-golang/services"
	"user-management-golang/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	service services.AuthService
}

func NewAuthHandler(s services.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.BindJSON(&req); err != nil {
		utils.JSONError(c, http.StatusBadRequest, "Invalid data")
		return
	}

	response, err := h.service.Login(req.Username, req.Password)
	if err != nil {
		utils.JSONError(c, http.StatusNotFound, err.Error())
		return
	}
	utils.JSONSuccess(c, response, "Login success")
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req model.RegisterRequest
	if err := c.BindJSON(&req); err != nil {
		utils.JSONError(c, http.StatusBadRequest, "Invalid data")
		return
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Name":
				errors = append(errors, "Nama wajib diisi dan minimal 3 karakter")
			case "Username":
				errors = append(errors, "Username wajib diisi dan minimal 3 karakter")
			case "Age":
				errors = append(errors, "Umur harus antara 18 sampai 60")
			case "Password":
				errors = append(errors, "Password minimal 8 character")
			default:
				errors = append(errors, err.Error())
			}
		}
		utils.JSONError(c, http.StatusBadRequest, errors)
		return
	}

	user, err := h.service.Register(utils.ToEntity(req))
	if err != nil {
		utils.JSONError(c, http.StatusBadRequest, err.Error())
		return
	}

	utils.JSONSuccess(c, utils.ToResponse(*user), "Register success")
}
