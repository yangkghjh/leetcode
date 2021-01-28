package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFirstMissingPositive(t *testing.T) {
	Convey("[41] 缺失的第一个正数", t, func() {
		args := [][]int{
			[]int{1, 2, 3, 4, 5, 6},
			[]int{-1, -2, 1, 3, 5, 6},
			[]int{-1, -2},
			[]int{1, 2, 0},
			[]int{3, 4, -1, 1},
			[]int{5, 8, 9, 11, 12},
			[]int{},
			[]int{1},
			[]int{2},
			[]int{1, 1},
			[]int{2, 2, 2},
			[]int{-2, -1, -3},
			[]int{2, 1},
		}
		expected := []int{7, 2, 1, 3, 2, 1, 1, 2, 1, 2, 1, 1, 3}

		for i, arg := range args {
			So(firstMissingPositive(arg), ShouldEqual, expected[i])
		}
	})
}
