package router

import (
	"user-management-golang/controller"
	"user-management-golang/middleware"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	r.GET("/users", middleware.AuthMiddleware(), controller.GetUser)
	r.POST("/users", controller.Register)
	r.GET("/users/:id", controller.GetDetailUser)
	r.POST("/users/login", controller.Login)
}
