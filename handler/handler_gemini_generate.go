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

	res, err := h.generateContent()

	if err != nil {
		return views.GenerateApiResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
	} else {
		return views.GenerateApiResponse(ctx, http.StatusOK, "", res)
	}
}
