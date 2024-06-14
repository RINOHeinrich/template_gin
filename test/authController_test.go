package controllers_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/RINOBE/gestion_de_livre/controllers"
	"github.com/RINOBE/gestion_de_livre/models"
)

func TestSignup(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Set up a test route for the Signup handler
	router.POST("/signup", controllers.Signup())
	phone := "1234567890"
	email := "test@example.com"
	password := "password"
	first_name := "John"
	last_name := "Doe"
	user_type := "ADMIN"
	// Create a new user for testing
	user := models.User{
		Email:      &email,
		Password:   &password,
		First_name: &first_name,
		Last_name:  &last_name,
		User_type:  &user_type,
		Phone:      &phone,
	}

	// Convert the user struct to JSON
	payload, _ := json.Marshal(user)

	// Create a new HTTP request with the JSON payload
	req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(payload))

	// Set the request content type to JSON
	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP recorder to record the response
	rec := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(rec, req)

	// Check the response status code
	assert.Equal(t, http.StatusInternalServerError, rec.Code)

	// TODO: Add more assertions to validate the response body or perform additional checks
}
