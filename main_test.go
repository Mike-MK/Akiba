package main

import (
	"akiba/controllers"
	"akiba/models"
	"bytes"
	"encoding/json"

	// "io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	models.ConnectToDatabase()
	router := gin.Default()
	return router
}

func TestLoginController(t *testing.T) {
	// mockResponse := `{"message":"Welcome to the Tech Company listing API with Golang"}`
	r := SetUpRouter()
	r.POST("/api/login", controllers.Login)
	creds := make(map[string]string,2)
	creds["username"] = "username"
	creds["password"] = "password"
	jsonBody,_ := json.Marshal(creds)

	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonBody))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRegisterController(t *testing.T) {
	// mockResponse := `{"message":"Welcome to the Tech Company listing API with Golang"}`
	r := SetUpRouter()
	r.POST("/api/register", controllers.Register)
	creds := make(map[string]string,2)
	creds["username"] = "username"
	creds["password"] = "password"
	jsonBody,_ := json.Marshal(creds)

	req, _ := http.NewRequest("POST", "/api/register", bytes.NewBuffer(jsonBody))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}