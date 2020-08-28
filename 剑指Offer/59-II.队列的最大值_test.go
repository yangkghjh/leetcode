package offer

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxQueue(t *testing.T) {
	Convey("59-II.队列的最大值", t, func() {
		obj := NewMaxQueue()
		So(obj.Max_value(), ShouldEqual, -1)
		obj.Push_back(3)
		So(obj.Max_value(), ShouldEqual, 3)
		obj.Push_back(4)
		So(obj.Max_value(), ShouldEqual, 4)
		obj.Push_back(2)
		So(obj.Max_value(), ShouldEqual, 4)
		So(obj.Pop_front(), ShouldEqual, 3)
		So(obj.Pop_front(), ShouldEqual, 4)
		So(obj.Pop_front(), ShouldEqual, 2)
		So(obj.Pop_front(), ShouldEqual, -1)
	})
}
