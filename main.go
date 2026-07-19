package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Define a counter metric: counts total HTTP requests, labelled by path
var requestCount = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "app_requests_total",
		Help: "Total number of HTTP requests, by path",
	},
	[]string{"path"},
)

func main() {
	// Main page - increments the counter for "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestCount.WithLabelValues("/").Inc()
		fmt.Fprintln(w, "Hello from the pipeline app!")
	})

	// Health check - increments the counter for "/health"
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		requestCount.WithLabelValues("/health").Inc()
		fmt.Fprintln(w, "OK")
	})

	// The metrics endpoint Prometheus will scrape
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}