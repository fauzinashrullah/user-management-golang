package controller

import (
	"net/http"
	"strconv"
	"user-management-golang/config"
	"user-management-golang/db"
	"user-management-golang/model"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Register(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		jsonError(c, http.StatusBadRequest, "Invalid data")
		return
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
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
		jsonError(c, http.StatusBadRequest, errors)
		return
	}

	if err := db.Database().Where("username = ?", user.Username).First(&user).Error; err == nil {
		jsonError(c, http.StatusBadRequest, "Username not available")
		return
	}

	if hash, err := config.HashPassword(user.Password); err != nil {
		jsonError(c, http.StatusBadRequest, "Invalid data")
		return
	} else {
		user.Password = hash
	}

	db.Database().Create(&user)
	response := toResponse(user)
	jsonSuccess(c, response, "Register success")
}

func Login(c *gin.Context) {
	var req loginRequest
	if err := c.BindJSON(&req); err != nil {
		jsonError(c, http.StatusBadRequest, "Invalid data")
		return
	}

	var user model.User
	if err := db.Database().Where("username = ?", req.Username).First(&user).Error; err != nil {
		jsonError(c, http.StatusNotFound, "User not found")
		return
	}

	if !config.VerifyPassword(req.Password, user.Password) {
		jsonError(c, http.StatusBadRequest, "Invalid password")
		return
	}

	token, _ := config.GenerateJwt(strconv.FormatUint(uint64(user.ID), 10))
	tokenResponse := map[string]string{"Token": token}

	response := toResponse(user)

	var responses []any
	responses = append(responses, response)
	responses = append(responses, tokenResponse)
	jsonSuccess(c, responses, "Login success")
}
