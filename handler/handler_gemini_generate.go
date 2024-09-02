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

func NewGenerateHandler(e *echo.Group, service *service.GenerateServices) *GenerateHandler {
	handler := GenerateHandler{service: service}

	e.POST("/generate", handler.generateContent)

	return &handler
}

func (h GenerateHandler) generateContent(ctx echo.Context) error {

	res := h.service.GenerateContent(ctx)

	if res != nil {
		return views.GenerateApiResponse(ctx, http.StatusInternalServerError, "Not generated", nil)
	} else {
		return views.GenerateApiResponse(ctx, http.StatusOK, "", res)
	}
}
