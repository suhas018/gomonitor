package storage

import (
	"sync"
	"time"

	"github.com/suhas018/gomonitor/internal/types"
)

type MemoryStorage struct {
	mu      sync.RWMutex
	metrics map[string][]types.Metric
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		metrics: make(map[string][]types.Metric),
	}
}

func (s *MemoryStorage) Store(m types.Metric) {
	s.mu.Lock()
	defer s.mu.Unlock()

	cutoff := time.Now().Add(-24 * time.Hour).Unix()
	metrics := s.metrics[m.Name]

	for i, metric := range metrics {
		if metric.Timestamp > cutoff {
			metrics = metrics[i:]
			break
		}
	}

	metrics = append(metrics, m)
	s.metrics[m.Name] = metrics
}

func (s *MemoryStorage) Query(name string, start, end int64) []types.Metric {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []types.Metric
	metrics := s.metrics[name]
	for _, m := range metrics {
		if m.Timestamp >= start && m.Timestamp <= end {
			result = append(result, m)
		}
	}
	return result
}

func (s *MemoryStorage) ListMetrics() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()

	names := make([]string, 0, len(s.metrics))
	for name := range s.metrics {
		names = append(names, name)
	}
	return names
}
