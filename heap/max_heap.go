package heap

import "sync"

func NewMaxHeap(capacity int) Heap {
	return &maxHeap{
		slice: make(slice, 0, capacity),
	}
}

type maxHeap struct {
	mu    sync.RWMutex
	slice slice
}

func (mh *maxHeap) Pop() (Item, error) {
	mh.mu.Lock()
	defer mh.mu.Unlock()

	if mh.Len() < 1 {
		return nil, ErrHeapEmpty
	}

	pop := mh.slice[0]
	mh.slice[0] = mh.slice[mh.Len()-1]
	mh.slice.DeleteLast()

	mh.shiftDown(0)

	return pop, nil
}

func (mh *maxHeap) Push(item Item) {
	mh.mu.Lock()
	defer mh.mu.Unlock()

	mh.slice = append(mh.slice, item)
	mh.shiftUp(mh.Len() - 1)
}

func (mh *maxHeap) IsEmpty() bool {
	mh.mu.RLock()
	defer mh.mu.RUnlock()
	return mh.Len() == 0
}

func (mh *maxHeap) Len() int {
	return len(mh.slice)
}

func (mh *maxHeap) shiftUp(index int) {
	parent := (index - 1) / 2
	for index > 0 && mh.slice[index].Compare(mh.slice[parent]) > 0 {
		mh.slice.Swap(index, parent)
		index = parent
		parent = (index - 1) / 2
	}
}

func (mh *maxHeap) shiftDown(index int) {
	for left, ok := mh.slice.LeftChild(index); ok; left, ok = mh.slice.LeftChild(index) {
		k := left
		if right, ok := mh.slice.RightChild(index); ok && mh.slice[left].Compare(mh.slice[right]) < 0 {
			k = right
		}

		if mh.slice[index].Compare(mh.slice[k]) > 0 {
			break
		}

		mh.slice.Swap(index, k)
		index = k
	}
}
