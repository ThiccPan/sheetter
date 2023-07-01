package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/thiccpan/sheetter/internal/usecase"
)

func NewUserRoute(router *echo.Echo, uu usecase.UserUsecase) {
	userHandler := NewSheetHandler(uu)
	router.GET("/users/data", userHandler.GetSheetData)
}
