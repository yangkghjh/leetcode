package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestClimbStairs(t *testing.T) {
	Convey("[69] x 的平方根", t, func() {
		args := []int{1, 2, 3, 4, 44}
		expected := []int{1, 2, 3, 5, 1134903170}

		for i, arg := range args {
			So(climbStairs(arg), ShouldResemble, expected[i])
		}
	})
}
