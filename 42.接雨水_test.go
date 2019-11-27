package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTrap(t *testing.T) {
	Convey("[42] 接雨水", t, func() {
		args := [][]int{
			[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			[]int{0},
			[]int{1, 2},
			[]int{1, 2, 1, 2, 1},
		}
		expected := []int{6, 0, 0, 1}

		for i, arg := range args {
			So(trap(arg), ShouldEqual, expected[i])
		}
	})
}
