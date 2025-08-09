package controller

import (
	"strconv"
	"strings"

	"user-management-golang/db"
	"user-management-golang/model"
	"user-management-golang/security"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type loginRequest struct {
	Name     string
	Password string
}
type userResponse struct {
	ID   uint
	Name string
	Age  int
}

func GetUser(c *gin.Context) {
	name := strings.ToLower(c.Query("name"))
	var users []model.User

	query := db.Database()
	if name != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+name+"%")
	}

	query.Find(&users)
	var responses []userResponse
	for _, user := range users {
		response := toResponse(user)
		responses = append(responses, response)
	}
	c.JSON(200, responses)
}

func Register(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "invalid data"})
		return
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Name":
				errors = append(errors, "Nama wajib diisi dan minimal 3 karakter")
			case "Age":
				errors = append(errors, "Umur harus antara 18 sampai 60")
			case "Password":
				errors = append(errors, "Password minimal 8 character")
			default:
				errors = append(errors, err.Error())
			}
		}
		c.JSON(400, gin.H{"error": errors})
		return
	}

	if hash, err := security.HashPassword(user.Password); err != nil {
		c.JSON(400, gin.H{"error": "invalid data"})
		return
	} else {
		user.Password = hash
	}

	db.Database().Create(&user)
	response := toResponse(user)
	c.JSON(200, response)
}

func Login(c *gin.Context) {
	var req loginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	var user model.User
	if err := db.Database().Where("name = ?", req.Name).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "Not found"})
		return
	}

	if !security.VerifyPassword(req.Password, user.Password) {
		c.JSON(400, gin.H{"error": "Password not valid"})
		return
	}

	token, _ := security.GenerateJwt(strconv.FormatUint(uint64(user.ID), 10))
	tokenResponse := map[string]string{"Token": token}

	response := toResponse(user)

	var responses []any
	responses = append(responses, response)
	responses = append(responses, tokenResponse)
	c.JSON(200, responses)
}

func GetDetailUser(c *gin.Context) {
	id := c.Param("id")
	var user model.User
	if err := db.Database().First(&user, "id = ?", id).Error; err != nil {
		c.JSON(400, gin.H{"error": "not found"})
		return
	}

	response := toResponse(user)

	c.JSON(200, response)
}

func toResponse(user model.User) userResponse {
	var response userResponse
	response.ID = user.ID
	response.Name = user.Name
	response.Age = user.Age
	return response
}
