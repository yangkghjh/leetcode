package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNumDecodings(t *testing.T) {
	Convey("[91] 解码方法", t, func() {
		args := []string{
			"12",
			"226",
			"101",
			"10",
			"0",
			"27",
			"100",
			"10001",
			"230",
			"301",
			"17",
		}
		expected := []int{
			2,
			3,
			1,
			1,
			0,
			1,
			0,
			0,
			0,
			0,
			2,
		}

		for i, arg := range args {
			So(numDecodings(arg), ShouldResemble, expected[i])
		}
	})
}
