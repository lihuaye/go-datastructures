package heap

import "errors"

var (
	ErrHeapEmpty = errors.New("the heap is empty")
)

type Item interface {
	Compare(other Item) int
}

type slice []Item

func (s *slice) DeleteLast() {
	size := len(*s)
	*s = (*s)[:size-1]
}

func (s *slice) Swap(dst, src int) {
	(*s)[src], (*s)[dst] = (*s)[dst], (*s)[src]
}

func (s *slice) LeftChild(index int) (int, bool) {
	left := index*2 + 1
	return left, left < len(*s)
}

func (s *slice) RightChild(index int) (int, bool) {
	right := index*2 + 2
	return right, right < len(*s)
}

type Heap interface {
	Pop() (Item, error)
	Push(Item)
	IsEmpty() bool
	Len() int
}
