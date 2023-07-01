package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thiccpan/sheetter/internal/usecase"
)

type UserHandler interface {
	GetSheetData(c echo.Context) error
	CreateSheetData(c echo.Context) error
}

type userHandler struct {
	uu usecase.UserUsecase
}

func NewSheetHandler(u usecase.UserUsecase) UserHandler {
	return &userHandler{
		uu: u,
	}
}

func (uh *userHandler) GetSheetData(c echo.Context) error {
	data, err := uh.uu.GetAllData()
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError, 
			map[string]interface{}{
				"error": err.Error(),
			})
	}

	return c.JSON(
		http.StatusInternalServerError, 
		map[string]interface{}{
			"data": data,
		})
}

func (uh *userHandler) CreateSheetData(c echo.Context) error {
	return c.JSON(http.StatusOK, "data")
}
