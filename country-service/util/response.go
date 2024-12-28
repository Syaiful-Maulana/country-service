package util

import (
	"github.com/labstack/echo/v4"
)

type Meta struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Response[T any] struct {
	Meta Meta `json:"meta"`
	Data T    `json:"data"`
}

func NewResponse[T any](c echo.Context, statusCode int, statusMessage string, message string, data T) error {
	meta := Meta{
		Code:    statusCode,
		Status:  statusMessage,
		Message: message,
	}

	return c.JSON(statusCode, Response[T]{
		Meta: meta,
		Data: data,
	})
}
