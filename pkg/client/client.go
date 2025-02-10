package client

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/suhas018/gomonitor/internal/types"
)

type Client struct {
	mu      sync.RWMutex
	metrics map[string]*types.Metric
}

func NewClient() *Client {
	return &Client{
		metrics: make(map[string]*types.Metric),
	}
}

func (c *Client) Counter(name string, labels map[string]string) *types.Metric {
	c.mu.Lock()
	defer c.mu.Unlock()

	if m, ok := c.metrics[name]; ok {
		return m
	}

	m := &types.Metric{
		Name:   name,
		Labels: labels,
		Value:  0,
	}
	c.metrics[name] = m
	return m
}

func (c *Client) Gauge(name string, value float64, labels map[string]string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.metrics[name] = &types.Metric{
		Name:   name,
		Value:  value,
		Labels: labels,
	}
}

func (c *Client) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	metrics := make([]types.Metric, 0, len(c.metrics))
	for _, m := range c.metrics {
		metric := *m
		metric.Timestamp = time.Now().Unix()
		metrics = append(metrics, metric)
	}

	json.NewEncoder(w).Encode(metrics)
}

func (c *Client) StartServer(addr string) error {
	http.Handle("/metrics", c)
	return http.ListenAndServe(addr, nil)
}
