package middleware

import (
	"fulka-api/util"
	"net/http"
	"time"
)

func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		util.HttpRequests.WithLabelValues(r.Method, r.URL.Path).Inc()

		duration := time.Since(start).Seconds()
		util.RequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)
	})
}
