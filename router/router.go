package router

import (
	"user-management-golang/db"
	"user-management-golang/handlers"
	"user-management-golang/middleware"
	"user-management-golang/services"

	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	r := gin.Default()

	db := db.Database()

	authService := services.NewAuthService(db)
	authHandler := handlers.NewAuthHandler(authService)

	userService := services.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService)

	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)
	r.GET("/users", middleware.AuthMiddleware(), userHandler.CurrentUser)
	r.GET("/admin/users", middleware.AuthMiddleware(), middleware.RequireRole("admin"), userHandler.GetUser)
	r.GET("/admin/users/:id", middleware.AuthMiddleware(), middleware.RequireRole("admin"), userHandler.GetUserDetail)

	return r
}
