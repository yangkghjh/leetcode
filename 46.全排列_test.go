package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPermute(t *testing.T) {
	Convey("[46] 全排列", t, func() {
		args := [][]int{
			[]int{1},
			[]int{1, 2},
			[]int{-1, -2, -3},
		}
		expected := [][][]int{
			[][]int{[]int{1}},
			[][]int{
				[]int{1, 2},
				[]int{2, 1},
			},
			[][]int{
				[]int{-1, -2, -3},
				[]int{-1, -3, -2},
				[]int{-2, -1, -3},
				[]int{-2, -3, -1},
				[]int{-3, -2, -1},
				[]int{-3, -1, -2},
			},
		}

		for i, arg := range args {
			So(permute(arg), ShouldResemble, expected[i])
		}
	})
}
