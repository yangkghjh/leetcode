package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMySqrt(t *testing.T) {
	Convey("[69] x 的平方根", t, func() {
		args := []int{0, 1, 2, 3, 4, 100, 8, 1111}
		expected := []int{0, 1, 1, 1, 2, 10, 2, 33}

		for i, arg := range args {
			So(mySqrt(arg), ShouldResemble, expected[i])
		}
	})
}
