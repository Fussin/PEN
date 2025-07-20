package queue

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
)

type RedisQueue struct {
	client *redis.Client
	ctx    context.Context
	key    string
}

func NewRedisQueue(addr, key string) *RedisQueue {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return &RedisQueue{
		client: rdb,
		ctx:    context.Background(),
		key:    key,
	}
}

func (q *RedisQueue) Enqueue(job Job) error {
	payload, err := json.Marshal(job)
	if err != nil {
		return err
	}
	return q.client.LPush(q.ctx, q.key, payload).Err()
}

func (q *RedisQueue) Dequeue() (*Job, error) {
	result, err := q.client.BRPop(q.ctx, 0, q.key).Result()
	if err != nil {
		return nil, err
	}

	var job Job
	err = json.Unmarshal([]byte(result[1]), &job)
	if err != nil {
		return nil, err
	}

	return &job, nil
}
