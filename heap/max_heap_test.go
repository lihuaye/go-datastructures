package heap

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type TestItem int

func (tt TestItem) Compare(item Item) int {
	if i, ok := item.(TestItem); ok {
		if tt < i {
			return -1
		} else if tt == i {
			return 0
		} else {
			return 1
		}
	}

	return -1
}

func TestMaxHeap_Push(t *testing.T) {
	heap := NewMaxHeap(10)

	Convey("Test push function", t, func() {
		Convey("Push [1,5,2,3,8] should slice be [8,5,2,1,3]", func() {
			heap.Push(TestItem(1))
			heap.Push(TestItem(5))
			heap.Push(TestItem(2))
			heap.Push(TestItem(3))
			heap.Push(TestItem(8))

			mh := heap.(*maxHeap)
			So(fmt.Sprint(mh.slice), ShouldEqual, fmt.Sprint([]int{8, 5, 2, 1, 3}))
		})
	})

}

func TestMaxHeap_IsEmpty(t *testing.T) {
	heap := NewMaxHeap(10)
	Convey("Test isEmpty function", t, func() {
		Convey("IsEmpty should return true that initially MaxHeap", func() {
			So(heap.IsEmpty(), ShouldBeTrue)
		})

		Convey("IsEmpty should return false that add item into MaxHeap", func() {
			heap.Push(TestItem(1))
			So(heap.IsEmpty(), ShouldBeFalse)
		})
	})
}

func TestMaxHeap_Pop(t *testing.T) {
	heap := NewMaxHeap(10)
	heap.Push(TestItem(1))
	heap.Push(TestItem(5))
	heap.Push(TestItem(2))
	heap.Push(TestItem(3))
	heap.Push(TestItem(8))

	Convey("Test pop function", t, func() {
		Convey("1. pop should returns 8,nil", func() {
			item, err := heap.Pop()
			So(err, ShouldBeNil)
			So(int(item.(TestItem)), ShouldEqual, 8)
		})

		Convey("2. pop should returns 5,nil", func() {
			item, err := heap.Pop()
			So(err, ShouldBeNil)
			So(int(item.(TestItem)), ShouldEqual, 5)
		})

		Convey("3. pop should returns 3,nil", func() {
			item, err := heap.Pop()
			So(err, ShouldBeNil)
			So(int(item.(TestItem)), ShouldEqual, 3)
		})

		Convey("4. pop should returns 2,nil", func() {
			item, err := heap.Pop()
			So(err, ShouldBeNil)
			So(int(item.(TestItem)), ShouldEqual, 2)
		})

		Convey("5. pop should returns 1,nil", func() {
			item, err := heap.Pop()
			So(err, ShouldBeNil)
			So(int(item.(TestItem)), ShouldEqual, 1)
		})

		Convey("6. pop should returns nil,err", func() {
			_, err := heap.Pop()
			So(err, ShouldBeError)
		})
	})
}

func TestMaxHeap_Len(t *testing.T) {
	heap := NewMaxHeap(10)
	Convey("Test len function", t, func() {
		Convey("MaxHeap_Len should return 0 when initially MaxHeap", func() {
			size := heap.Len()
			So(size, ShouldEqual, 0)
		})

		Convey("Add items into heap", func() {
			heap.Push(TestItem(1))
			heap.Push(TestItem(2))
			heap.Push(TestItem(3))
			Convey("Len should return 3 when added items", func() {
				size := heap.Len()
				So(size, ShouldEqual, 3)
			})
		})

		Convey("Pop items from heap", func() {
			_, _ = heap.Pop()
			Convey("Len should return 2 when pop one item", func() {
				size := heap.Len()
				So(size, ShouldEqual, 2)
			})

			_, _ = heap.Pop()
			_, _ = heap.Pop()
			_, _ = heap.Pop()
			Convey("Len should return 0 when pop all item", func() {
				size := heap.Len()
				So(size, ShouldEqual, 0)
			})
		})
	})
}
