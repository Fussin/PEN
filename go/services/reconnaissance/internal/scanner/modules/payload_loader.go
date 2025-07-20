package modules

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
)

type PayloadManager struct {
	mu              sync.Mutex
	customPayloads  []Payload
	contextPayloads map[string][]Payload
}

func NewPayloadManager() *PayloadManager {
	return &PayloadManager{
		contextPayloads: make(map[string][]Payload),
	}
}

func (pm *PayloadManager) LoadPayloadsFromFile(filePath string) error {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var payloads []Payload
	err = json.Unmarshal(byteValue, &payloads)
	if err != nil {
		return err
	}

	pm.mu.Lock()
	defer pm.mu.Unlock()
	pm.customPayloads = append(pm.customPayloads, payloads...)
	return nil
}

func (pm *PayloadManager) GetPayloadsForContext(ctx string) []Payload {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	return pm.contextPayloads[ctx]
}
