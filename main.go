package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func main() {
	r := gin.Default()

	r.Run(":8080")
	fmt.Println("hello")
}
