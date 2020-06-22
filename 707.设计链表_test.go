package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConstructor(t *testing.T) {
	Convey("[54] 螺旋矩阵", t, func() {
		linkedList := &MyLinkedList{}
		linkedList.AddAtHead(1)
		linkedList.DeleteAtIndex(0)
		linkedList.AddAtHead(1)
		linkedList.AddAtTail(3)
		linkedList.AddAtIndex(1, 2)           //链表变为1-> 2-> 3
		So(linkedList.Get(1), ShouldEqual, 2) //返回2
		linkedList.DeleteAtIndex(1)           //现在链表是1-> 3
		So(linkedList.Get(1), ShouldEqual, 3) //返回3
		linkedList.DeleteAtIndex(10)
		linkedList.AddAtTail(3)
		linkedList.AddAtIndex(10, 2)
		linkedList.Get(10)
	})
}
