package api

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/suhas018/gomonitor/internal/types"
)

type Server struct {
	storage types.Storage
	addr    string
	router  *mux.Router
}

func NewServer(storage types.Storage, addr string) *Server {
	s := &Server{
		storage: storage,
		addr:    addr,
		router:  mux.NewRouter(),
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.router.HandleFunc("/metrics", s.handleMetricsList).Methods("GET")
	s.router.HandleFunc("/query", s.handleQuery).Methods("GET")
}

func (s *Server) Start() error {
	return http.ListenAndServe(s.addr, s.router)
}
