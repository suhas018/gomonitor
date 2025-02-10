package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/suhas018/gomonitor/internal/api"
	"github.com/suhas018/gomonitor/internal/scraper"
	"github.com/suhas018/gomonitor/internal/storage"
)

func main() {
	log.Println("Starting monitoring system...")

	// Initialize components
	store := storage.NewMemoryStorage()

	// Modified to only scrape from the example service
	scraper := scraper.NewScraper(store, []string{
		"http://localhost:8081/metrics", // Only scrape from the example service
	})

	apiServer := api.NewServer(store, ":8080")

	// Start services
	go func() {
		log.Printf("Starting scraper...")
		scraper.Start()
	}()

	go func() {
		log.Printf("Starting API server on :8080...")
		if err := apiServer.Start(); err != nil {
			log.Fatalf("API server error: %v", err)
		}
	}()

	// Wait for shutdown signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("Shutting down...")
}
