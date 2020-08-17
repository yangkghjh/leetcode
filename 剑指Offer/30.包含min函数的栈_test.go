package offer

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMinStack(t *testing.T) {
	Convey("最小栈", t, func() {
		minStack := NewMinStack()

		minStack.Push(-2)
		minStack.Push(0)
		minStack.Push(-3)
		So(minStack.Min(), ShouldEqual, -3) // 返回 -3.
		minStack.Pop()
		So(minStack.Top(), ShouldEqual, 0)  // 返回 0.
		So(minStack.Min(), ShouldEqual, -2) // 返回 -2.
	})
}
