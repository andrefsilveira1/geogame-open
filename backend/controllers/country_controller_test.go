package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	db "geolocation/database"
	"geolocation/models"
	"geolocation/repositories"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetCountry(t *testing.T) {
	app := fiber.New()

	// Set up a test database connection
	db, err := db.ConnectAzure()
	assert.NoError(t, err)
	defer db.Close()

	// Create a test repository
	repository := repositories.NewCountryRepo(db)
	fmt.Println(repository)

	// Create a test Fiber app instance
	app.Get("/country/:id", func(c *fiber.Ctx) error {
		// Call the controller function
		err := GetCountry(c)
		return err
	})

	// Create a test request
	req := httptest.NewRequest(http.MethodGet, "/country/1", nil)
	res, err := app.Test(req, -1) // The second parameter -1 means "do not print Fiber logs"
	assert.NoError(t, err)
	defer res.Body.Close()

	// Verify the response status code
	assert.Equal(t, http.StatusOK, res.StatusCode)

	// Retrieve the response body
	var response models.Country
	err = json.NewDecoder(res.Body).Decode(&response)
	assert.NoError(t, err)

	// Verify the response structure and values
	assert.Equal(t, uint(1), response.Id)
	assert.NotEmpty(t, response.Name)
	assert.NotEmpty(t, response.Tips)
	assert.NotEmpty(t, response.Level)
	assert.NotEmpty(t, response.Score)
	assert.NotEmpty(t, response.Location)

	// You can add more assertions based on your specific expectations
}

func TestSendScore(t *testing.T) {
	// Initialize Fiber app
	app := fiber.New()

	// Set up a test database connection
	db, err := db.ConnectAzure()
	assert.NoError(t, err)
	defer db.Close()

	// Create a test repository
	repository := repositories.NewRepository(db)
	fmt.Println(repository)

	// Create a test Fiber app instance
	app.Post("/send-score", func(c *fiber.Ctx) error {
		// Call the controller function
		err := SendScore(c)
		return err
	})

	// Create a sample score data
	scoreData := models.Score{
		// Populate with sample data as needed
	}

	// Convert score data to JSON
	payload, err := json.Marshal(scoreData)
	assert.NoError(t, err)

	// Create a test request with JSON payload
	req := httptest.NewRequest(http.MethodPost, "/send-score", bytes.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")

	// Perform the request
	res, err := app.Test(req, -1) // The second parameter -1 means "do not print Fiber logs"
	assert.NoError(t, err)
	defer res.Body.Close()

	// Verify the response status code
	assert.Equal(t, http.StatusOK, res.StatusCode)

	// Retrieve the response body
	var response interface{}
	err = json.NewDecoder(res.Body).Decode(&response)
	assert.NoError(t, err)

}
