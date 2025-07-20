package queue

import "sync"

type MockQueue struct {
	jobs []Job
	mu   sync.Mutex
}

func NewMockQueue() *MockQueue {
	return &MockQueue{jobs: make([]Job, 0)}
}

func (q *MockQueue) Enqueue(job Job) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.jobs = append(q.jobs, job)
	return nil
}

func (q *MockQueue) Dequeue() (*Job, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.jobs) == 0 {
		return nil, nil // Or return an error to indicate an empty queue
	}
	job := q.jobs[0]
	q.jobs = q.jobs[1:]
	return &job, nil
}
