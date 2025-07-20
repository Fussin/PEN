package scanner

import (
	"github.com/autonomouspen/reconnaissance/pkg/plugin"
	"github.com/autonomouspen/reconnaissance/pkg/queue"
	"github.com/autonomouspen/reconnaissance/pkg/ratelimiter"
	"testing"
	"time"
)

type MockPlugin struct{}

func (p *MockPlugin) Name() string {
	return "Mock Plugin"
}

func (p *MockPlugin) Run(target string) (string, error) {
	return "mock result", nil
}

func TestScanner(t *testing.T) {
	pluginManager := plugin.NewManager()
	pluginManager.Register(&MockPlugin{})

	rateLimiter := ratelimiter.NewRateLimiter(100, time.Second)
	jobQueue := queue.NewMockQueue()

	s := NewScanner(1, pluginManager, rateLimiter, jobQueue)
	s.Start()

	jobQueue.Enqueue(queue.Job{ID: "test", Target: "test-target"})

	select {
	case result := <-s.Results():
		if result != "mock result" {
			t.Errorf("Expected 'mock result', got '%s'", result)
		}
	case <-time.After(5 * time.Second):
		t.Error("Timed out waiting for result")
	}
}
