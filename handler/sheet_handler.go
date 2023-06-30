package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type SheetHandler interface {
	GetSheetData(c echo.Context) error
	CreateSheetData(c echo.Context) error
}

type sheetHandler struct {
}

func NewSheetHandler() SheetHandler {
	return new(sheetHandler)
}

func (sh *sheetHandler) GetSheetData(c echo.Context) error {
	return c.JSON(http.StatusOK, "data")
}

func (sh *sheetHandler) CreateSheetData(c echo.Context) error {
	return c.JSON(http.StatusOK, "data")
}