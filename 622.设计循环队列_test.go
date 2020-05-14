package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMyCircularQueue(t *testing.T) {
	Convey("[622] 设计循环队列", t, func() {
		obj := Constructor(3)
		So(obj.DeQueue(), ShouldEqual, false)
		So(obj.Len(), ShouldEqual, 0)
		So(obj.EnQueue(1), ShouldEqual, true)
		So(obj.Len(), ShouldEqual, 1)
		So(obj.EnQueue(2), ShouldEqual, true)
		So(obj.Len(), ShouldEqual, 2)
		So(obj.EnQueue(3), ShouldEqual, true)
		So(obj.Len(), ShouldEqual, 3)
		So(obj.EnQueue(4), ShouldEqual, false)
		So(obj.Len(), ShouldEqual, 3)
		So(obj.Front(), ShouldEqual, 1)
		So(obj.Rear(), ShouldEqual, 3)
		So(obj.DeQueue(), ShouldBeTrue)
		So(obj.Front(), ShouldEqual, 2)
		So(obj.Len(), ShouldEqual, 2)
		So(obj.EnQueue(4), ShouldBeTrue)

		obj = Constructor(4)
		obj.EnQueue(3)
		obj.Front()
		obj.IsFull()
		obj.EnQueue(7)
		obj.EnQueue(2)
		obj.EnQueue(5)
		So(obj.Len(), ShouldEqual, 4)
		obj.DeQueue()
		So(obj.Len(), ShouldEqual, 3)
		So(obj.head, ShouldEqual, 1)
		So(obj.tail, ShouldEqual, 3)
		obj.EnQueue(4)
		So(obj.head, ShouldEqual, 1)
		So(obj.tail, ShouldEqual, 0)
		So(obj.Len(), ShouldEqual, 4)
		So(obj.EnQueue(2), ShouldBeFalse)
		obj.IsEmpty()
		obj.Rear()
	})
}
