package main

import (
	"github.com/didinj/go-jwt-auth/config"
	"github.com/didinj/go-jwt-auth/controllers"
	"github.com/didinj/go-jwt-auth/models"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{})

	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.POST("/refresh", controllers.RefreshToken)
	r.POST("/logout", controllers.AuthMiddleware(), controllers.Logout)

	protected := r.Group("/api")
	protected.Use(controllers.AuthMiddleware())
	protected.GET("/protected", controllers.Protected)

	r.Run(":8080")
}
