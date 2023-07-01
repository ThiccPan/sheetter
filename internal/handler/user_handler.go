package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/thiccpan/sheetter/internal/entity"
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
	data, err := uh.uu.GetAllData(c.Request().Context())
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]interface{}{
				"error": err.Error(),
			})
	}

	return c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"message": "get sheet data successfull",
			"data":    data,
		})
}

func (uh *userHandler) CreateSheetData(c echo.Context) error {
	request := UserReq{}
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]interface{}{
				"error": err.Error(),
			})
	}

	data, err := uh.uu.CreateData(
		c.Request().Context(),
		entity.User{
			Name:        request.Name,
			Email:       request.Email,
			IsContacted: false,
			Row:         request.Row,
		})

	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]interface{}{
				"error": err.Error(),
			})
	}

	return c.JSON(
		http.StatusOK,
		map[string]interface{}{
			"message": "data created successfully",
			"data":    data,
		})
}

type UserReq struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Row   int    `json:"row"`
}
