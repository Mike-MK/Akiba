package main

import (
	"akiba/controllers"
	"akiba/models"
	"akiba/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectToDatabase()
	r := gin.Default()
	
	// routes
	public := r.Group("/api")
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)
	
	protected := r.Group("/api/account")
	protected.Use(middleware.JWTMiddleware())
	protected.POST("/new",controllers.CreateAccount)


	r.Run("localhost:8080")
	fmt.Println("hello")
}
