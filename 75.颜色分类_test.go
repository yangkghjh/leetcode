package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSortColors(t *testing.T) {
	Convey("[73] 矩阵置零", t, func() {
		args := [][]int{
			[]int{2, 0, 2, 1, 1, 0},
		}
		expected := [][]int{
			[]int{0, 0, 1, 1, 2, 2},
		}

		for i, arg := range args {
			sortColors(arg)
			So(arg, ShouldResemble, expected[i])
		}
	})
}
