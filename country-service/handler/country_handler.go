package handler

import (
	"encoding/json"
	models "fulka-api/models/country"
	"fulka-api/service"
	"fulka-api/util"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type CountryHandler struct {
	countryService service.CountryService
}

func NewCountryHandler(countryService service.CountryService) *CountryHandler {
	return &CountryHandler{countryService: countryService}
}

func (h *CountryHandler) GetAllCountries(w http.ResponseWriter, r *http.Request) {
	log.Printf("GetAllCountries request received")

	query := r.URL.Query()
	page := query.Get("page")
	pageNumber := query.Get("pageNumber")

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
		util.WriteJSONResponse(w, http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}

	countCountries, err := h.countryService.CountAllCountry()
	if err != nil {
		util.WriteJSONResponse(w, http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}

	response := map[string]interface{}{
		"data":        countries,
		"CurrentPage": pageInt,
		"PerPage":     pageNumberInt,
		"Total":       countCountries,
	}
	util.WriteJSONResponse(w, http.StatusOK, "success", "all countries", response)
}

func (h *CountryHandler) GetByIdCountries(w http.ResponseWriter, r *http.Request) {
	log.Printf("GetByIdCountries request received")

	id := strings.TrimPrefix(r.URL.Path, "/countries/")

	country, err := h.countryService.GetByIdCountries(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.WriteJSONResponse(w, http.StatusOK, "success", "country", country)
}

func (h *CountryHandler) CreateCountry(w http.ResponseWriter, r *http.Request) {
	log.Printf("CreateCountry request received")

	var country models.Country
	if err := json.NewDecoder(r.Body).Decode(&country); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	err := h.countryService.CreateCountry(&country)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	util.WriteJSONResponse(w, http.StatusOK, "success", "country", country)
}
