package router

import (
	"user-management-golang/controller"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	r.GET("/users", controller.GetUser)
	r.POST("/users", controller.CreateUser)
	r.GET("/users/:id", controller.GetDetailUser)
}
