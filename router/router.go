package router

import (
	"user-management-golang/controller"
	"user-management-golang/middleware"

	"github.com/gin-gonic/gin"
)

func Route(r *gin.Engine) {
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	r.GET("/users", middleware.AuthMiddleware(), controller.GetUser)
	r.GET("/users/:id", middleware.AuthMiddleware(), controller.GetDetailUser)
}
