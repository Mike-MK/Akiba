package controllers

import (
	"akiba/models"
	"akiba/utils/mpesa"
	"akiba/utils/token"
	"bytes"
	b64 "encoding/base64"
	"fmt"
	// "io"
	"net/http"
	// "os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
)

type AccountInput struct{
	Number string `json:"number" binding:"required"`
}

type MpesaRequestBody struct{
	BusinessShortCode int 
	Password string
	Timestamp string
	TransactionType string
	Amount int
	PartyA int
	PartyB int
	PhoneNumber int
	CallBackURL string
	AccountReference string 
	TransactionDesc string
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

func Deposit(c *gin.Context){
	authToken,err := mpesa.GetMpesaAuthToken();
	fmt.Println("token",authToken)

	if  err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"message":"Error getting auth token"})
		return 
	}
	url := "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest"

	// stk push request body
	timestamp := time.Now().Format("20060102150405")
	fmt.Println(timestamp)
	code := 174379
	key := "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919"
	phone := 254716537782
	pwd := fmt.Sprint(code)+key+timestamp
	amount := 1

	b64Pwd := b64.StdEncoding.EncodeToString([]byte(pwd))
	// data := make(map[string]string,11)
	// data["BusinessShortCode"] = code
	// data["Password"] = b64Pwd
	// data["Timestamp"] = timestamp
	// data["TransactionType"] = "CustomerPayBillOnline"
	// data["Amount"] = amount
	// data["PartyA"] = phone
	// data["PartyB"] = code
	// data["PhoneNumber"] = phone
	// data["CallBackURL"] = "https://c2d5-102-140-246-229.ngrok.io/wallet/result/"
	// data["AccountReference"] = phone
	// data["TransactionDesc"] = "Akiba Pay"

	var input MpesaRequestBody
	input.AccountReference = fmt.Sprint(code)
	input.Amount = amount
	input.BusinessShortCode = code
	input.CallBackURL = "https://c2d5-102-140-246-229.ngrok.io/wallet/result/"
	input.PartyA = phone
	input.PartyB = code
	input.Password = b64Pwd
	input.PhoneNumber = phone
	input.Timestamp = timestamp
	input.TransactionDesc = "Akiba Pay"
	input.TransactionType = "CustomerPayBillOnline"

	jsonBody,err := json.Marshal(input)
	if err != nil{
		fmt.Println(err)
	}

	bodyReader := bytes.NewReader(jsonBody)
	fmt.Println(bodyReader)

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"message":"Error"})
		return 
	}
	authHeader := "Bearer "+ authToken
	req.Header.Add("Authorization", authHeader)
	req.Header.Add("Content-Type", "application/json")
	

	resp, err := client.Do(req)
	if err != nil {
		return 
	}
	var respBody map[string]string
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	
	c.JSON(http.StatusOK,gin.H{"message":"Ok"})
}
