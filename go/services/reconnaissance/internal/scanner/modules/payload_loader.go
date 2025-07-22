package modules

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	"github.com/autonomouspen/reconnaissance/internal/common"
)

type PayloadManager struct {
	mu              sync.Mutex
	contextPayloads map[string][]common.Payload
}

func NewPayloadManager() *PayloadManager {
	return &PayloadManager{
		contextPayloads: make(map[string][]common.Payload),
	}
}

func (pm *PayloadManager) LoadPayloadsFromFile(filePath string) error {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var payloads []common.Payload
	err = json.Unmarshal(byteValue, &payloads)
	if err != nil {
		return err
	}

	pm.mu.Lock()
	defer pm.mu.Unlock()
	for _, p := range payloads {
		pm.contextPayloads[p.Context] = append(pm.contextPayloads[p.Context], p)
	}
	return nil
}

func (pm *PayloadManager) GetPayloadsForContext(ctx string) []common.Payload {
	pm.mu.Lock()
	defer pm.mu.Unlock()
	return pm.contextPayloads[ctx]
}
