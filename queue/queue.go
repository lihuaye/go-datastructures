package queue

import "errors"

var (
	ErrQueueEmpty = errors.New("queue is empty")
)

type Queue interface {
	Enqueue(interface{})
	Dequeue() (interface{}, error)
	Len() int
}
