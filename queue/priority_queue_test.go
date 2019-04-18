package queue

import (
	"github.com/lihuaye/go-datastructures/heap"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type TestItem int

func (tt TestItem) Compare(item heap.Item) int {
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

func TestPriorityQueue(t *testing.T) {
	queue := NewPriorityQueue(10)
	Convey("Test Priority Queue", t, func() {
		Convey("Len returns should be 0 that priority queue is empty", func() {
			So(queue.Len(), ShouldBeZeroValue)
		})

		Convey("Len returns should be 6 that add six of items into the queue", func() {
			queue.Enqueue(TestItem(40))
			queue.Enqueue(TestItem(20))
			queue.Enqueue(TestItem(60))
			queue.Enqueue(TestItem(10))
			queue.Enqueue(TestItem(70))
			queue.Enqueue(TestItem(50))
			So(queue.Len(), ShouldEqual, 6)
		})

		Convey("Pop Test", func() {
			Convey("1. Pop returns should be 70, nil", func() {
				item, err := queue.Dequeue()
				So(err, ShouldBeNil)
				So(int(item.(TestItem)), ShouldEqual, 70)
			})

			Convey("2. Pop returns should be 100 that add a item of TestItem(100)", func() {
				queue.Enqueue(TestItem(100))
				item, err := queue.Dequeue()
				So(err, ShouldBeNil)
				So(int(item.(TestItem)), ShouldEqual, 100)
			})

			Convey("3. Pop returns should be 60", func() {
				item, err := queue.Dequeue()
				So(err, ShouldBeNil)
				So(int(item.(TestItem)), ShouldEqual, 60)
			})

			Convey("4. Pop returns should be 50", func() {
				item, err := queue.Dequeue()
				So(err, ShouldBeNil)
				So(int(item.(TestItem)), ShouldEqual, 50)
			})

			Convey("5. Pop returns should be 40", func() {
				item, err := queue.Dequeue()
				So(err, ShouldBeNil)
				So(int(item.(TestItem)), ShouldEqual, 40)
			})

			Convey("6. Pop returns should be 20", func() {
				item, err := queue.Dequeue()
				So(err, ShouldBeNil)
				So(int(item.(TestItem)), ShouldEqual, 20)
			})

			Convey("7. Pop returns should be 10", func() {
				item, err := queue.Dequeue()
				So(err, ShouldBeNil)
				So(int(item.(TestItem)), ShouldEqual, 10)
			})

			Convey("8. Pop returns should be 0,error", func() {
				_, err := queue.Dequeue()
				So(err, ShouldBeError)
			})
		})
	})
}
