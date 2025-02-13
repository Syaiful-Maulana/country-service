package routes

import (
	"database/sql"
	"fulka-api/handler"
	"fulka-api/middleware"
	"fulka-api/repository"
	"fulka-api/service"
	"net/http"
)

func CountryRoutes(mux *http.ServeMux, db *sql.DB) {
	if db == nil {
		panic("Database connection is nil")
	}

	countryRepo := repository.NewCountryRepository(db)
	countryService := service.NewCountryService(countryRepo, db)
	countryHandler := handler.NewCountryHandler(countryService)

	mux.Handle("/countries", middleware.PrometheusMiddleware(middleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			countryHandler.GetAllCountries(w, r)
		case http.MethodPost:
			countryHandler.CreateCountry(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))))

	mux.Handle("/countries/", middleware.PrometheusMiddleware(middleware.JWTMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			countryHandler.GetByIdCountries(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}))))
}
