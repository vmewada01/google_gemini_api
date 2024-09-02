package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type GenerateServices struct{}

func NewGenerateServices() *GenerateServices {
	return &GenerateServices{}
}

func (g *GenerateServices) GenerateContent(c echo.Context) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("GEMINI_KEY")

	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash-latest:generateContent?key=" + apiKey

	body := map[string]interface{}{
		"contents": []map[string]interface{}{
			{
				"parts": []map[string]interface{}{
					{"text": "Explain how AI works"},
				},
			},
		},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error encoding JSON",
		})
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error creating request",
		})
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error making the request",
		})
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Error reading response",
		})
	}

	return c.JSONBlob(http.StatusOK, respBody)
}
