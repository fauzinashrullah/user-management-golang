package controller

import (
	"net/http"
	"strings"

	"user-management-golang/db"
	"user-management-golang/model"

	"github.com/gin-gonic/gin"
)

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
	jsonSuccess(c, responses, "Get user success")
}

func GetDetailUser(c *gin.Context) {
	id := c.Param("id")
	var user model.User
	if err := db.Database().First(&user, "id = ?", id).Error; err != nil {
		jsonError(c, http.StatusNotFound, "User not found")
		return
	}

	response := toResponse(user)

	jsonSuccess(c, response, "Get detail user success")
}
