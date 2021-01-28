package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSearchRange(t *testing.T) {
	Convey("[34] 在排序数组中查找元素的第一个和最后一个位置", t, func() {
		args1 := [][]int{
			{},
			{1},
			{1, 4},
			{1, 2, 3, 3, 3, 4, 5},
			{1, 1, 3, 3, 3, 4, 5},
			{1, 1, 3, 3, 3, 5, 5},
			{5, 7, 7, 8, 8, 10},
			{1, 1, 1, 3, 3, 5, 5},
		}
		args2 := []int{0, 1, 4, 3, 1, 5, 6, 3}
		expected := [][]int{
			{-1, -1},
			{0, 0},
			{1, 1},
			{2, 4},
			{0, 1},
			{5, 6},
			{-1, -1},
			{3, 4},
		}

		for k, v := range args1 {
			So(searchRange(v, args2[k]), ShouldResemble, expected[k])
		}
	})
}
