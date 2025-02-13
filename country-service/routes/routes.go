package routes

import (
	"database/sql"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux, db *sql.DB) {
	CountryRoutes(mux, db)
	LoggerRoutes(mux, db)
}
