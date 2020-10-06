package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSubsets(t *testing.T) {
	Convey("[78] 子集", t, func() {
		args := [][]int{
			[]int{1, 2, 3},
		}
		expected := []int{
			8,
		}

		for i, arg := range args {
			So(len(subsets(arg)), ShouldResemble, expected[i])
		}
	})
}
