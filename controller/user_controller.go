package controller

import (
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
	c.JSON(200, users)
}

func CreateUser(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "invalid data"})
		return
	}

	db.Database().Create(&user)
	c.JSON(200, user)
}
func GetDetailUser(c *gin.Context) {
	id := c.Param("id")
	var user model.User
	if err := db.Database().First(&user, "id = ?", id).Error; err != nil {
		c.JSON(400, gin.H{"error": "not found"})
		return
	}

	c.JSON(200, user)
}
