package handler

import (
	"go_gemini/service"
	"go_gemini/views"
	"net/http"

	"github.com/labstack/echo/v4"
)

type GenerateHandler struct {
	service *service.GenerateServices
}

func NewAffiliateHandler(e *echo.Group, service *service.GenerateServices) *GenerateHandler {
	handler := GenerateHandler{service: service}

	e.GET("/generate/", handler.generateContent)

	return &handler
}

func (h GenerateHandler) generateContent(ctx echo.Context) error {

	res, err := h.service.generateContent()

	if err != nil {
		return views.GenerateApiResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
	} else {
		return views.GenerateApiResponse(ctx, http.StatusOK, "", res)
	}
}

// func generateContent(c echo.Context) error {
// 	// Define the API key and URL
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	// Access the environment variable
// 	apiKey := os.Getenv("GEMINI_KEY")

// 	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash-latest:generateContent?key=" + apiKey

// 	// Create the request body
// 	body := map[string]interface{}{
// 		"contents": []map[string]interface{}{
// 			{
// 				"parts": []map[string]interface{}{
// 					{"text": "Explain how AI works"},
// 				},
// 			},
// 		},
// 	}

// 	// Convert the body to JSON
// 	jsonBody, err := json.Marshal(body)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{
// 			"error": "Error encoding JSON",
// 		})
// 	}

// 	// Create the request
// 	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{
// 			"error": "Error creating request",
// 		})
// 	}

// 	// Set the headers
// 	req.Header.Set("Content-Type", "application/json")

// 	// Send the request
// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{
// 			"error": "Error making the request",
// 		})
// 	}
// 	defer resp.Body.Close()

// 	// Read the response body
// 	respBody, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{
// 			"error": "Error reading response",
// 		})
// 	}

// 	// Return the response body as JSON
// 	return c.JSONBlob(http.StatusOK, respBody)
// }
