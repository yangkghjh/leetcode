package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestExist(t *testing.T) {
	Convey("[79] 单词搜索", t, func() {
		board := [][]byte{
			[]byte{'A', 'B', 'C', 'E'},
			[]byte{'S', 'F', 'C', 'S'},
			[]byte{'A', 'D', 'E', 'E'},
		}
		args := []string{
			"ABCCED",
			"SEE",
			"ABCB",
		}
		expected := []bool{
			true, true, false,
		}

		for i, arg := range args {
			So(exist(board, arg), ShouldResemble, expected[i])
		}
	})
}
