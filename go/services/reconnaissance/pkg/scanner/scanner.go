package scanner

import (
	"github.com/autonomouspen/reconnaissance/pkg/plugin"
	"github.com/autonomouspen/reconnaissance/pkg/queue"
	"github.com/autonomouspen/reconnaissance/pkg/ratelimiter"
	"log"
	"sync"
)

type Scanner struct {
	maxConcurrency int
	results        chan string
	wg             sync.WaitGroup
	pluginManager  *plugin.Manager
	rateLimiter    *ratelimiter.RateLimiter
	jobQueue       queue.Queue
}

func NewScanner(maxConcurrency int, pluginManager *plugin.Manager, rateLimiter *ratelimiter.RateLimiter, jobQueue queue.Queue) *Scanner {
	return &Scanner{
		maxConcurrency: maxConcurrency,
		results:        make(chan string),
		pluginManager:  pluginManager,
		rateLimiter:    rateLimiter,
		jobQueue:       jobQueue,
	}
}

func (s *Scanner) Start() {
	for i := 0; i < s.maxConcurrency; i++ {
		s.wg.Add(1)
		go s.worker()
	}

	go func() {
		s.wg.Wait()
		close(s.results)
	}()
}

func (s *Scanner) worker() {
	defer s.wg.Done()
	for {
		job, err := s.jobQueue.Dequeue()
		if err != nil {
			log.Printf("Error dequeuing job: %v", err)
			continue
		}
		if job == nil {
			continue
		}

		s.rateLimiter.Wait()
		results := s.pluginManager.Run(job.Target)
		for _, result := range results {
			s.results <- result
		}
	}
}

func (s *Scanner) Results() <-chan string {
	return s.results
}
