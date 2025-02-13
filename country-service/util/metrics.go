package util

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var HttpRequests = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total jumlah request HTTP per endpoint",
	},
	[]string{"method", "endpoint"},
)

var RequestDuration = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Histogram durasi request HTTP dalam detik",
		Buckets: prometheus.DefBuckets,
	},
	[]string{"method", "endpoint"},
)

func InitMetrics() {
	prometheus.MustRegister(HttpRequests, RequestDuration)
}
