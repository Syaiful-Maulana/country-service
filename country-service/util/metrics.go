package util

import (
	"runtime"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Total jumlah request HTTP per endpoint
var HttpRequests = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total jumlah request HTTP per endpoint",
	},
	[]string{"method", "endpoint"},
)

// Histogram durasi request HTTP dalam detik
var RequestDuration = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Histogram durasi request HTTP dalam detik",
		Buckets: prometheus.DefBuckets,
	},
	[]string{"method", "endpoint"},
)

// Total jumlah error HTTP berdasarkan status code
var HttpErrors = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_errors_total",
		Help: "Total jumlah error HTTP per endpoint",
	},
	[]string{"method", "endpoint", "status_code"},
)

// Jumlah request HTTP yang sedang diproses
var ActiveRequests = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "http_active_requests",
		Help: "Jumlah request HTTP yang sedang diproses",
	},
	[]string{"method", "endpoint"},
)

// Histogram ukuran response HTTP dalam byte
var ResponseSize = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_response_size_bytes",
		Help:    "Histogram ukuran response HTTP dalam byte",
		Buckets: prometheus.ExponentialBuckets(100, 2, 10), // Mulai dari 100 byte dengan kelipatan 2
	},
	[]string{"method", "endpoint"},
)

// Histogram durasi query database dalam detik
var DBQueryDuration = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "db_query_duration_seconds",
		Help:    "Histogram durasi query database dalam detik",
		Buckets: prometheus.DefBuckets,
	},
	[]string{"operation", "table"},
)

// Jumlah goroutine yang sedang berjalan
var CustomGoroutines = promauto.NewGauge(
	prometheus.GaugeOpts{
		Name: "custom_go_goroutines",
		Help: "Jumlah goroutine yang sedang berjalan (custom metric)",
	},
)

// Penggunaan memori aplikasi dalam byte
var MemoryUsage = promauto.NewGauge(
	prometheus.GaugeOpts{
		Name: "go_memory_usage_bytes",
		Help: "Penggunaan memori aplikasi dalam byte",
	},
)

// Fungsi untuk memperbarui metrik penggunaan memori dan goroutine
func UpdateRuntimeMetrics() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	MemoryUsage.Set(float64(memStats.Alloc))
	CustomGoroutines.Set(float64(runtime.NumGoroutine()))
}

// Registrasi semua metrik
func InitMetrics() {
	prometheus.MustRegister(
		HttpRequests,
		RequestDuration,
		HttpErrors,
		ActiveRequests,
		ResponseSize,
		DBQueryDuration,
		CustomGoroutines,
		MemoryUsage,
	)
}
