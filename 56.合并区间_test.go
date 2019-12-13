package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMerge(t *testing.T) {
	Convey("[56] 合并区间", t, func() {
		args := [][][]int{
			[][]int{
				[]int{1, 3},
				[]int{2, 6},
				[]int{8, 10},
				[]int{15, 18},
			},
			[][]int{
				[]int{1, 4},
				[]int{4, 5},
			},
			[][]int{
				[]int{1, 4},
			},
			[][]int{},
		}
		expected := [][][]int{
			[][]int{
				[]int{1, 6},
				[]int{8, 10},
				[]int{15, 18},
			},
			[][]int{
				[]int{1, 5},
			},
			[][]int{
				[]int{1, 4},
			},
			[][]int{},
		}

		for i, arg := range args {
			So(merge(arg), ShouldResemble, expected[i])
		}
	})
}
