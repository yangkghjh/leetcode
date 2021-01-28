package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCountAndSay(t *testing.T) {
	Convey("[38] 报数", t, func() {
		args := []int{1, 2, 3, 4, 5, 6}
		expected := []string{"1", "11", "21", "1211", "111221", "312211"}

		for i, arg := range args {
			So(countAndSay(arg), ShouldEqual, expected[i])
		}
	})
}
