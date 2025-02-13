package routes

import (
	"database/sql"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func LoggerRoutes(mux *http.ServeMux, db *sql.DB) {
	mux.Handle("/metrics", promhttp.Handler())
}
