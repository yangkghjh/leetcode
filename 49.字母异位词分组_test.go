package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGroupAnagrams(t *testing.T) {
	Convey("[49] 字母异位词分组", t, func() {
		args := [][]string{
			[]string{},
			[]string{"123"},
			[]string{"123", "231", "234"},
		}
		expected := [][][]string{
			[][]string{},
			[][]string{
				[]string{"123"},
			},
			[][]string{
				[]string{"123", "231"},
				[]string{"234"},
			},
		}

		for i, arg := range args {
			result := groupAnagrams(arg)
			So(len(result), ShouldEqual, len(expected[i]))
			for _, e := range expected[i] {
				So(result, ShouldContain, e)
			}
		}

		So(groupAnagramsSort("bca"), ShouldEqual, "abc")
	})
}
