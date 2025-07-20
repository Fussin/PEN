package queue

type Job struct {
	ID     string
	Target string
}

type Queue interface {
	Enqueue(job Job) error
	Dequeue() (*Job, error)
}
