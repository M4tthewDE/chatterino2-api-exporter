package server

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(port string) *Server {
	mux := http.NewServeMux()
	httpServer := &http.Server{Addr: ":" + port, Handler: mux}

	s := &Server{
		httpServer: httpServer,
	}

	mux.Handle("/metrics", promhttp.Handler())
	return s
}

func (s *Server) ListenAndServe() {
	log.Println("Starting HTTP server")

	err := s.httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
