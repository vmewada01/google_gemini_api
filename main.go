package main

import (
	"go_gemini/handler"
	"go_gemini/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	apiGroup := e.Group("/api")

	// Define the API route
	geminiGroup := apiGroup.Group("/gemini")
	geminiService := &service.GenerateServices{}
	handler.NewGenerateHandler(geminiGroup, geminiService)

	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
