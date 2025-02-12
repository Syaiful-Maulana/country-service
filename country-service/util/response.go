package util

import (
	"encoding/json"
	"net/http"

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

func WriteJSONResponse(w http.ResponseWriter, status int, result, message string, data interface{}) {
	meta := Meta{
		Code:    status,
		Status:  result,
		Message: message,
	}
	response := map[string]interface{}{
		"data": data,
		"meta": meta,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}
