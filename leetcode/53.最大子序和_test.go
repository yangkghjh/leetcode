package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMaxSubArray(t *testing.T) {
	Convey("[53] 最大子序和", t, func() {
		args := [][]int{
			[]int{1, 2, 3},
			[]int{1},
			[]int{},
			[]int{2, -1, 3},
			[]int{-1, -2, -3},
			[]int{-5, -2, -1, -7},
			[]int{-5, -2, -1, -7, 100},
			[]int{-5, -2, -1, -7, 5, 100},
			[]int{-5, -2, 5, -1, -7, -10, 100},
			[]int{-5, -2, 5, -1, -7, -10},
			[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			[]int{-1, 1, 2, 1},
			[]int{-2, 3, 1, 3},
		}
		expected := []int{
			6, 1, 0, 4, -1, -1, 100, 105, 100, 5, 6, 4, 7,
		}

		for i, arg := range args {
			So(maxSubArray(arg), ShouldResemble, expected[i])
		}
	})
}
