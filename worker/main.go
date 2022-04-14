package worker

import (
	"errors"
	"fmt"
	"github.com/samdevbr/gandalf/queue"
	"log"
	"os"
	"sync"
)

type worker struct {
	q *queue.Queue
}

func (w *worker) Start(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		filename, err := w.q.Pop()

		if errors.Is(err, queue.ErrQueueIsEmpty) {
			break
		}

		source, err := os.ReadFile(filename)

		if err != nil {
			log.Println(err)
			continue
		}

		fmt.Println(len(source))
	}
}

func New(q *queue.Queue) *worker {
	return &worker{
		q,
	}
}
