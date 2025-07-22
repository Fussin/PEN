package payloads

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/autonomouspen/scanner/internal/common"
)

type Manager struct {
	mu       sync.Mutex
	payloads []common.Payload
}

func NewManager(payloadsFile string) *Manager {
	manager := &Manager{}
	if err := manager.loadPayloads(payloadsFile); err != nil {
		log.Printf("Error loading payloads: %v", err)
	}
	return manager
}

func (m *Manager) loadPayloads(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	m.mu.Lock()
	defer m.mu.Unlock()
	return json.Unmarshal(bytes, &m.payloads)
}

func (m *Manager) GetPayloads() []common.Payload {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.payloads
}

type Payload struct {
	Value   string `json:"Value"`
	Context string `json:"Context"`
}
