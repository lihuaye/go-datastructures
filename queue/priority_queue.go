package queue

import (
	"errors"
	"github.com/lihuaye/go-datastructures/heap"
)

type priorityQueue struct {
	heap heap.Heap
}

func NewPriorityQueue(capacity int) Queue {
	return &priorityQueue{
		heap: heap.NewMaxHeap(capacity),
	}
}

func (pq *priorityQueue) Enqueue(item interface{}) {
	if i, ok := item.(heap.Item); ok {
		pq.heap.Push(i)
	} else {
		panic(errors.New("item not implements the heap.Item interface"))
	}
}

func (pq *priorityQueue) Dequeue() (interface{}, error) {
	item, err := pq.heap.Pop()
	if err == heap.ErrHeapEmpty {
		return nil, ErrQueueEmpty
	}
	return item, nil
}

func (pq *priorityQueue) Len() int {
	return pq.heap.Len()
}
