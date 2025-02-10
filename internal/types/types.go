package types

// Metric represents a single data point in time series
type Metric struct {
	Name      string            `json:"name"`
	Value     float64           `json:"value"`
	Labels    map[string]string `json:"labels"`
	Timestamp int64             `json:"timestamp"`
}

// Storage defines the interface for metric storage implementations
type Storage interface {
	Store(Metric)
	Query(string, int64, int64) []Metric
	ListMetrics() []string
}
