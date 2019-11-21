package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSearch(t *testing.T) {
	Convey("[34] 在排序数组中查找元素的第一个和最后一个位置", t, func() {
		args1 := [][]int{
			[]int{},
			[]int{1},
			[]int{1, 4},
			[]int{1, 2, 3, 3, 3, 4, 5},
			[]int{1, 1, 3, 3, 3, 4, 5},
			[]int{1, 1, 3, 3, 3, 5, 5},
		}
		args2 := []int{0, 1, 4, 3, 1, 5}
		expected := [][]int{
			[]int{-1, -1},
			[]int{0, 0},
			[]int{1, 1},
			[]int{2, 4},
			[]int{0, 1},
			[]int{5, 6},
		}

		for k, v := range args1 {
			So(searchRange(v, args2[k]), ShouldResemble, expected[k])
		}
	})
}
