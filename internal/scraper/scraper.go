package scraper

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/suhas018/gomonitor/internal/types"
)

type Scraper struct {
	storage  types.Storage
	targets  []string
	client   *http.Client
	interval time.Duration
}

func NewScraper(storage types.Storage, targets []string) *Scraper {
	return &Scraper{
		storage:  storage,
		targets:  targets,
		client:   &http.Client{Timeout: 10 * time.Second},
		interval: 15 * time.Second,
	}
}

func (s *Scraper) Start() {
	ticker := time.NewTicker(s.interval)
	go func() {
		for range ticker.C {
			s.scrape()
		}
	}()
}

func (s *Scraper) scrape() {
	for _, target := range s.targets {
		resp, err := s.client.Get(target)
		if err != nil {
			log.Printf("Error scraping %s: %v", target, err)
			continue
		}

		var metrics []types.Metric
		if err := json.NewDecoder(resp.Body).Decode(&metrics); err != nil {
			log.Printf("Error decoding metrics from %s: %v", target, err)
			resp.Body.Close()
			continue
		}
		resp.Body.Close()

		for _, m := range metrics {
			m.Timestamp = time.Now().Unix()
			s.storage.Store(m)
		}
	}
}
