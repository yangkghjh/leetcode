package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSetZeroes(t *testing.T) {
	Convey("[73] 矩阵置零", t, func() {
		args := [][][]int{
			[][]int{
				[]int{1, 1, 1},
				[]int{1, 0, 1},
				[]int{1, 1, 1},
			},
		}
		expected := [][][]int{
			[][]int{
				[]int{1, 0, 1},
				[]int{0, 0, 0},
				[]int{1, 0, 1},
			},
		}

		for i, arg := range args {
			setZeroes(arg)
			So(arg, ShouldResemble, expected[i])
		}
	})
}
