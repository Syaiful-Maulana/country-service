package routes

import (
	"database/sql"
	"fulka-api/middleware"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterRoutes(mux *http.ServeMux, db *sql.DB) {
	mux.Handle("/metrics", promhttp.Handler())

	CountryRoutes(mux, db)

	mux.Handle("/", middleware.PrometheusMiddleware(middleware.NotFoundMiddleware(mux)))
}
