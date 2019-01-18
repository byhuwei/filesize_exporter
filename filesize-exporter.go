package main

import (
"log"
"net/http"

"filesize_exporter/collector"
"github.com/prometheus/client_golang/prometheus"
"github.com/prometheus/client_golang/prometheus/promhttp"
)



func main() {
        var test string = "test"
	metrics := collector.NewMetrics(test)
	registry := prometheus.NewRegistry()
	registry.MustRegister(metrics)

	http.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))

	log.Fatal(http.ListenAndServe(":9090", nil))
}

