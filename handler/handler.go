package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Server struct {
	sync.Mutex
	names map[string]int
}

func NewServer() *Server {
	return &Server{names: make(map[string]int)}
}

// HelloHandler responds with Hello, <name>!.
func (s *Server) HelloHandler(w http.ResponseWriter, r *http.Request) {
	s.Lock()
	defer s.Unlock()

	name := r.URL.Path[len("/hello/"):]

	s.names[name]++

	fmt.Fprintf(w, "Hello, %s!", name)
}

// OKHandler responds with OK
func OKHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

// CountsHandler GET returns each name has been called; DELETE resets count
func (s *Server) CountsHandler(w http.ResponseWriter, r *http.Request) {
	s.Lock()
	defer s.Unlock()

	if r.Method == http.MethodDelete {
		s.names = map[string]int{}

		w.WriteHeader(http.StatusNoContent)
	}

	type data struct {
		Name  string `json:"name"`
		Count int    `json:"count"`
	}

	resp := make([]data, 0)
	for name, count := range s.names {
		resp = append(resp, data{Name: name, Count: count})
	}

	if err := json.NewEncoder(w).Encode(&resp); err != nil {
		log.Println("Encoding error:", err)
		w.WriteHeader(501)
	}
}

// HealthHandler measures time elapsed since start of service
func HealthHandler(version string) http.HandlerFunc {
	startTime := time.Now()

	return func(w http.ResponseWriter, r *http.Request) {
		h := struct {
			Uptime  string `json:"uptime"`
			Status  string `json:"status"`
			Version string `json:"version"`
		}{
			Status:  "OK",
			Uptime:  time.Since(startTime).String(),
			Version: version,
		}
		if err := json.NewEncoder(w).Encode(&h); err != nil {
			log.Println("Encoding error:", err)
			w.WriteHeader(501)
		}
	}
}
