package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCanJump(t *testing.T) {
	Convey("[55] 跳跃游戏", t, func() {
		args := [][]int{
			[]int{2, 3, 1, 1, 4},
			[]int{3, 2, 1, 0, 4},
			[]int{},
			[]int{0},
			[]int{1, 0},
			[]int{1, 1, 1, 1, 1},
		}
		expected := []bool{
			true,
			false,
			false,
			true,
			true,
			true,
		}

		for i, arg := range args {
			So(canJump(arg), ShouldResemble, expected[i])
		}
	})
}
