package main

import (
	"bytes"
	"encoding/json"
	"go_gemini/handler"
	"go_gemini/service"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Create a new Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	apiGroup := e.Group("/api")

	// Define the API route
	GeminiGroup := apiGroup.Group("/gemini")
	GeminiService := service.GenerateServices()
	handler.GenerateHandler(GeminiGroup, GeminiService)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler function for the /generate route
func generateContent(c echo.Context) error {
	// Define the API key and URL
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Access the environment variable
	apiKey := os.Getenv("GEMINI_KEY")

	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash-latest:generateContent?key=" + apiKey

	// Create the request body
	body := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]interface{}{
					{"text": "Explain how AI works"},
				},
			},
		},
	}

	// Convert the body to JSON
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error encoding JSON",
		})
	}

	// Create the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error creating request",
		})
	}

	// Set the headers
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error making the request",
		})
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error reading response",
		})
	}

	// Return the response body as JSON
	return c.JSONBlob(http.StatusOK, respBody)
}
