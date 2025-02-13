package middleware

import (
	"encoding/json"
	"net/http"
)

func NotFoundMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		response := map[string]string{
			"error":   "Not Found",
			"message": "The requested resource could not be found.",
		}

		json.NewEncoder(w).Encode(response)
	})
}
