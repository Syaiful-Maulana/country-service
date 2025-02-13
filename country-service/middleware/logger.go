package middleware

import (
	"fulka-api/util"
	"net/http"
	"strconv"
	"time"
)

type responseWriterRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (r *responseWriterRecorder) WriteHeader(code int) {
	r.statusCode = code
	r.ResponseWriter.WriteHeader(code)
}

func PrometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		recorder := &responseWriterRecorder{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(recorder, r)

		util.HttpRequests.WithLabelValues(r.Method, r.URL.Path).Inc()

		duration := time.Since(start).Seconds()
		util.RequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)

		if recorder.statusCode >= 400 {
			util.HttpErrors.WithLabelValues(r.Method, r.URL.Path, strconv.Itoa(recorder.statusCode)).Inc()
		}
	})
}
