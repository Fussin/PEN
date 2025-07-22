package modules

import (
	"github.com/autonomouspen/reconnaissance/internal/common"
	"sync"
)

type ResultCollector struct {
	mu              sync.Mutex
	vulnerabilities []common.Vulnerability
	screenshots     map[string][]byte
}

func NewResultCollector() *ResultCollector {
	return &ResultCollector{
		vulnerabilities: []common.Vulnerability{},
		screenshots:     make(map[string][]byte),
	}
}

func (r *ResultCollector) Add(v common.Vulnerability) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.vulnerabilities = append(r.vulnerabilities, v)
}

func (r *ResultCollector) AttachScreenshot(url string, image []byte) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.screenshots[url] = image
}

func (r *ResultCollector) All() []common.Vulnerability {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.vulnerabilities
}

func (r *ResultCollector) Screenshots() map[string][]byte {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.screenshots
}
