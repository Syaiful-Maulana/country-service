package handler

import (
	"log"
	"net/http"
	"strconv"

	"fulka-api/service"
	"fulka-api/util"

	"github.com/labstack/echo/v4"
)

type CountryHandler struct {
	countryService service.CountryService
}

func NewCountryHandler(countryService service.CountryService) *CountryHandler {
	return &CountryHandler{countryService: countryService}
}

func (h *CountryHandler) GetAllCountries(c echo.Context) error {
	log.Printf("GetAllCountries request received")

	page := c.QueryParam("page")
	pageNumber := c.QueryParam("pageNumber")

	if page == "" {
		page = "1"
	}

	if pageNumber == "" {
		pageNumber = "10"
	}

	pageInt, _ := strconv.Atoi(page)
	pageNumberInt, _ := strconv.Atoi(pageNumber)

	countries, err := h.countryService.GetAllCountries(pageInt, pageNumberInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	countCountries, err := h.countryService.CountAllCountry()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return util.NewResponse(c, http.StatusOK, "success", "all countries", map[string]interface{}{
		"data":        countries,
		"CurrentPage": pageInt,
		"PerPage":     pageNumberInt,
		"Total":       countCountries,
	})
}

func (h *CountryHandler) GetByIdCountries(c echo.Context) error {
	log.Printf("GetByIdCountries request received")

	id := c.Param("id")

	country, err := h.countryService.GetByIdCountries(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return util.NewResponse(c, http.StatusOK, "success", "country", country)
}
