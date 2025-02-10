package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func (s *Server) handleMetricsList(w http.ResponseWriter, r *http.Request) {
	metrics := s.storage.ListMetrics()
	json.NewEncoder(w).Encode(map[string]interface{}{
		"metrics": metrics,
	})
}

func (s *Server) handleQuery(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "metric name required", http.StatusBadRequest)
		return
	}

	startStr := r.URL.Query().Get("start")
	endStr := r.URL.Query().Get("end")

	var start, end int64
	var err error

	if startStr != "" {
		start, err = strconv.ParseInt(startStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid start time", http.StatusBadRequest)
			return
		}
	} else {
		start = time.Now().Add(-1 * time.Hour).Unix()
	}

	if endStr != "" {
		end, err = strconv.ParseInt(endStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid end time", http.StatusBadRequest)
			return
		}
	} else {
		end = time.Now().Unix()
	}

	metrics := s.storage.Query(name, start, end)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"metrics": metrics,
	})
}
