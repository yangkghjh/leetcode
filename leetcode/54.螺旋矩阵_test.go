package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpiralOrder(t *testing.T) {
	Convey("[54] 螺旋矩阵", t, func() {
		args := [][][]int{
			[][]int{
				[]int{1, 2, 3},
				[]int{4, 5, 6},
				[]int{7, 8, 9},
			},
			[][]int{
				[]int{1, 2, 3, 4},
				[]int{5, 6, 7, 8},
				[]int{9, 10, 11, 12},
			},
			[][]int{},
			[][]int{
				[]int{1, 2, 3},
			},
		}
		expected := [][]int{
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
			[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
			[]int{},
			[]int{1, 2, 3},
		}

		for i, arg := range args {
			So(spiralOrder(arg), ShouldResemble, expected[i])
		}
	})
}
