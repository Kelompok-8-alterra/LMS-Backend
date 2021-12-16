package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Meta struct {
		Status  int    `json:"status"`
		Message []string `json:"message"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func NewSuccesResponse(c echo.Context, data interface{}) error {
	response := BaseResponse{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = []string{"Success"}
	response.Data = data
	return c.JSON(http.StatusOK, response)
}

func NewErrorResponse(c echo.Context, status int, err error) error {
	response := BaseResponse{}
	response.Meta.Status = status
	response.Meta.Message = []string{err.Error()}
	response.Data = nil
	return c.JSON(status, response)
}