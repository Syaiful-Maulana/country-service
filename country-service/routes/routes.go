package routes

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo, db *sql.DB) {
	CountryRoutes(e, db)
}
