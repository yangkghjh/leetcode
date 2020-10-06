package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUntilEnd(t *testing.T) {
	Convey("[29] 两数相除", t, func() {
		So(divide(1, 2), ShouldEqual, 0)
		So(divide(2, 1), ShouldEqual, 2)
		So(divide(-2, 1), ShouldEqual, -2)
		So(divide(1, -1), ShouldEqual, -1)
		So(divide(45, 5), ShouldEqual, 9)
		So(divide(-2147483648, -1), ShouldEqual, 2147483647)
	})
}
