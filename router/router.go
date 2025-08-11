package router

import (
	"time"
	"user-management-golang/db"
	"user-management-golang/handlers"
	"user-management-golang/middleware"
	"user-management-golang/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Route() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // atau "*" untuk semua
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	db := db.Database()

	authService := services.NewAuthService(db)
	authHandler := handlers.NewAuthHandler(authService)

	userService := services.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService)

	r.POST("/api/register", authHandler.Register)
	r.POST("/api/login", authHandler.Login)
	r.GET("/api/users", middleware.AuthMiddleware(), userHandler.CurrentUser)
	r.GET("/api/admin/users", middleware.AuthMiddleware(), middleware.RequireRole("admin"), userHandler.GetUser)
	r.GET("/api/admin/users/:id", middleware.AuthMiddleware(), middleware.RequireRole("admin"), userHandler.GetUserDetail)

	return r
}
