package controllers

import (
	"akiba/models"
	"akiba/utils/token"
	"net/http"
	"github.com/gin-gonic/gin"
)

type AccountInput struct{
	Number string `json:"number" binding:"required"`
}

func CreateAccount(c *gin.Context){
	var input AccountInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	a := models.Account{}
	a.Number = input.Number
	
	
	user_id, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u,err := models.GetUserByID(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	a.User = u
	_,err = a.SaveAccount()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	c.JSON(http.StatusOK,gin.H{"message":"Authorized to view"})
}
