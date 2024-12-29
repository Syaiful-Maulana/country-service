package routes

import (
	"database/sql"
	"fulka-api/handler"
	"fulka-api/repository"
	"fulka-api/service"

	"github.com/labstack/echo/v4"
)

func CountryRoutes(e *echo.Echo, db *sql.DB) {
	if db == nil {
		panic("Database connection is nil")
	}

	countryRepo := repository.NewCountryRepository(db)
	countryService := service.NewCountryService(countryRepo, db)
	countryHandler := handler.NewCountryHandler(countryService)

	e.GET("/countries", countryHandler.GetAllCountries)
	e.GET("/countries/:id", countryHandler.GetByIdCountries)
	e.POST("/countries", countryHandler.CreateCountry)
}
