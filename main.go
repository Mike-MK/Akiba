package main

import (
	"akiba/controllers"
	"akiba/models"
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
	

	r.Run("localhost:8080")
	fmt.Println("hello")
}
