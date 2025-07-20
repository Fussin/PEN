package modules

import (
	"fmt"
	"net/http"
	"sync"
)

type BlindXSSServer struct {
	mu        sync.Mutex
	callbacks map[string]string
}

func NewBlindXSSServer() *BlindXSSServer {
	return &BlindXSSServer{
		callbacks: make(map[string]string),
	}
}

func (s *BlindXSSServer) Start() {
	http.HandleFunc("/xss", func(w http.ResponseWriter, r *http.Request) {
		s.mu.Lock()
		defer s.mu.Unlock()
		s.callbacks[r.URL.Query().Get("id")] = r.URL.RawQuery
		fmt.Println("[!] Blind XSS Callback! Query:", r.URL.RawQuery)
	})
	fmt.Println("Listening for Blind XSS callbacks on :8081")
	go http.ListenAndServe(":8081", nil)
}

func (s *BlindXSSServer) CheckCallback(id string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	callback, ok := s.callbacks[id]
	return callback, ok
}
