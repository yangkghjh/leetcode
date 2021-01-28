package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUniquePaths(t *testing.T) {
	Convey("[62] 不同路径", t, func() {
		args := [][]int{
			[]int{7, 3},
			[]int{5, 4},
			[]int{3, 2},
			[]int{0, 0},
			[]int{1, 1},
			[]int{2, 2},
			[]int{10, 10},
		}
		expected := []int{
			28,
			35,
			3,
			0,
			1,
			2,
			48620,
		}

		for i, arg := range args {
			So(uniquePaths(arg[0], arg[1]), ShouldEqual, expected[i])
		}
	})
}
