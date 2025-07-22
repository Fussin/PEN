package modules

import "sync"

type ResultCollector struct {
	mu              sync.Mutex
	vulnerabilities []Vulnerability
	screenshots     map[string][]byte
}

func NewResultCollector() *ResultCollector {
	return &ResultCollector{
		vulnerabilities: []Vulnerability{},
		screenshots:     make(map[string][]byte),
	}
}

func (r *ResultCollector) Add(v Vulnerability) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.vulnerabilities = append(r.vulnerabilities, v)
}

func (r *ResultCollector) AttachScreenshot(url string, image []byte) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.screenshots[url] = image
}

func (r *ResultCollector) All() []Vulnerability {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.vulnerabilities
}
