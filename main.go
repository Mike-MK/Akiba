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

	public := r.Group("/api")
	public.POST("/register", controllers.Register)
	r.Run("localhost:8080")
	fmt.Println("hello")
}
