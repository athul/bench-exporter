package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
  // TODO:improve this whole thing
	benchCollector := newBenchMetrics()
	prometheus.MustRegister(benchCollector)
	http.Handle("/metrics", promhttp.Handler())
	log.Println(http.ListenAndServe(":9101", nil))
}
