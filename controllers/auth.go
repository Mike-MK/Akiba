package controllers

import (
	"akiba/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct{
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct{
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}


func Register(c *gin.Context) {

	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	_,err := u.SaveUser()

	if err != nil {

		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusOK,gin.H{"message":"registered successfully"})
}

func Login(c *gin.Context) {
	var input LoginInput
	err := c.ShouldBindJSON(&input); if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}
	u := models.User{}
	u.Username = input.Username
	u.Password = input.Password

	token, err := models.LoginCheck(u.Username,u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token":token})
}

func CreateAccount(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{"message":"Authorized to view"})
}
