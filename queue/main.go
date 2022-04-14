package queue

import (
	"errors"
	"sync"
)

var ErrQueueIsEmpty = errors.New("queue is empty")

type Queue struct {
	mutex sync.Mutex
	items []string
}

func (q *Queue) Push(item string) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.items = append(q.items, item)
}

func (q *Queue) Pop() (string, error) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if len(q.items) == 0 {
		return "", ErrQueueIsEmpty
	}

	item := q.items[0]
	q.items = q.items[1:]

	return item, nil
}

func (q *Queue) Len() int {
	return len(q.items)
}
