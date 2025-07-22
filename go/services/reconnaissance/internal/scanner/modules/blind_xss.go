package modules

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type BlindXSSCallback struct {
	ID        string
	Query     string
	UserAgent string
	Headers   http.Header
	Timestamp time.Time
}

type BlindXSSServer struct {
	mu        sync.Mutex
	callbacks map[string]BlindXSSCallback
}

func NewBlindXSSServer() *BlindXSSServer {
	return &BlindXSSServer{
		callbacks: make(map[string]BlindXSSCallback),
	}
}

func (s *BlindXSSServer) Start() {
	http.HandleFunc("/xss", func(w http.ResponseWriter, r *http.Request) {
		s.mu.Lock()
		defer s.mu.Unlock()
		id := r.URL.Query().Get("id")
		callback := BlindXSSCallback{
			ID:        id,
			Query:     r.URL.RawQuery,
			UserAgent: r.UserAgent(),
			Headers:   r.Header,
			Timestamp: time.Now(),
		}
		s.callbacks[id] = callback
		fmt.Println("[!] Blind XSS Callback! Query:", r.URL.RawQuery)
	})
	http.HandleFunc("/xss/results", func(w http.ResponseWriter, r *http.Request) {
		s.mu.Lock()
		defer s.mu.Unlock()
		json.NewEncoder(w).Encode(s.callbacks)
	})
	fmt.Println("Listening for Blind XSS callbacks on :8081")
	go http.ListenAndServe(":8081", nil)
}

func (s *BlindXSSServer) CheckCallback(id string) (BlindXSSCallback, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	callback, ok := s.callbacks[id]
	return callback, ok
}
