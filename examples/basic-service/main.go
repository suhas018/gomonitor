package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/suhas018/gomonitor/pkg/client"
)

func main() {
	// Create metrics client
	metrics := client.NewClient()

	// Create example metrics
	requestCounter := metrics.Counter("http_requests_total", map[string]string{
		"service": "example",
		"method":  "GET",
	})

	// Create a gauge for response time
	metrics.Gauge("response_time_ms", 0, map[string]string{
		"service": "example",
	})

	// Simulate metrics
	go func() {
		for {
			// Increment request counter
			requestCounter.Value++

			// Update response time gauge
			metrics.Gauge("response_time_ms",
				rand.Float64()*100, // Random response time 0-100ms
				map[string]string{"service": "example"},
			)

			time.Sleep(time.Second)
		}
	}()

	log.Printf("Starting example service on :8081...")
	log.Fatal(metrics.StartServer(":8081"))
}
