package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var benchDir string
var userCounts bool

func main() {
	benchPath := flag.String("bench", "/home/frappe/frappe-bench", "Path to Bench Directory")
	usersCounts := flag.Bool("users", false, "Perform SQL operations to fetch user counts")
	flag.Parse()
	benchDir = *benchPath
	userCounts = *usersCounts

	benchCollector := newBenchMetrics()
	prometheus.MustRegister(benchCollector)
	http.Handle("/metrics", promhttp.Handler())

	log.Printf("INFO: Starting http server - %s", "0.0.0.0:9101")
	if err := http.ListenAndServe(":9101", nil); err != nil {
		log.Fatalf("ERROR: Failed to start http server: %s", err)
	}
}
