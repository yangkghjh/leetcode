package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestPlusOne(t *testing.T) {
	Convey("[66] 加一", t, func() {
		args := [][]int{
			[]int{7, 3},
			[]int{0},
			[]int{},
			[]int{9},
			[]int{1, 9, 9},
		}
		expected := [][]int{
			[]int{7, 4},
			[]int{1},
			[]int{1},
			[]int{1, 0},
			[]int{2, 0, 0},
		}

		for i, arg := range args {
			So(plusOne(arg), ShouldResemble, expected[i])
		}
	})
}
