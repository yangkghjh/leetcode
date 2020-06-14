package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMovingAverage_Next(t *testing.T) {
	Convey("", t, func() {
		obj := MovingAverageConstructor(2)
		So(obj.Next(2), ShouldEqual, 2)
		So(obj.Next(0), ShouldEqual, 1)
		So(obj.Next(1), ShouldEqual, 0.5)
		obj = MovingAverageConstructor(0)
		So(obj.Next(2), ShouldEqual, 0)
	})
}
