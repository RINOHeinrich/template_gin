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

func TestInsertBook(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()
	// Set up a test route for the InsertBook handler
	router.POST("/insert-book", controllers.InsertBook())
	var title string = "Notre Dame de Paris"
	var author string = "Victor Hugo"
	// Create a new book for testing
	book := models.Book{
		Title:  &title,
		Author: &author,
	}

	// Convert the book struct to JSON
	payload, _ := json.Marshal(book)

	// Create a new HTTP request with the JSON payload
	req, _ := http.NewRequest("POST", "/insert-book", bytes.NewBuffer(payload))
	// Set the request content type to JSON
	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP recorder to record the response
	rec := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(rec, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, rec.Code)

	// TODO: Add more assertions to validate the response body or perform additional checks
}
func TestGetAllBooks(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()
	// Set up a test route for the GetAllBooks handler
	router.GET("/get-all-books", controllers.GetAllBooks())
	// Create a new HTTP request
	req, _ := http.NewRequest("GET", "/get-all-books", nil)
	// Create a new HTTP recorder to record the response
	rec := httptest.NewRecorder()
	// Perform the request
	router.ServeHTTP(rec, req)
	// Check the response status code
	assert.Equal(t, http.StatusOK, rec.Code)
	// TODO: Add more assertions to validate the response body or perform additional checks
}
