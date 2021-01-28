package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSearch(t *testing.T) {
	Convey("[33] 搜索旋转排序数组", t, func() {
		So(searchRotateSortedArray([]int{1, 2, 3, 4, 5}, 2), ShouldEqual, 1)
		So(searchRotateSortedArray([]int{1, 2, 3, 4, 5}, 4), ShouldEqual, 3)
		So(searchRotateSortedArray([]int{1, 2, 3, 4, 5}, 5), ShouldEqual, 4)
		So(searchRotateSortedArray([]int{4, 5, 1, 2, 3}, 5), ShouldEqual, 1)
		So(searchRotateSortedArray([]int{2, 3, 4, 5, 1}, 3), ShouldEqual, 1)

		So(searchRotateSortedArray([]int{2, 3, 4, 5, 1}, 2), ShouldEqual, 0)
		So(searchRotateSortedArray([]int{2, 3, 4, 5, 1}, 1), ShouldEqual, 4)
		So(searchRotateSortedArray([]int{2, 3, 4, 5, 1}, 4), ShouldEqual, 2)
		So(searchRotateSortedArray([]int{2, 3, 4, 5, 1}, 7), ShouldEqual, -1)
		So(searchRotateSortedArray([]int{2, 3, 4, 6, 1}, 5), ShouldEqual, -1)
		So(searchRotateSortedArray([]int{}, 5), ShouldEqual, -1)
	})
}
