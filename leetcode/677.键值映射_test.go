package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMapSum(t *testing.T) {
	Convey("[677] 键值映射", t, func() {
		mapsum := &MapSum{
			Children: map[byte]*MapSum{},
			Map:      make(map[string]int),
		}

		mapsum.Insert("apple", 3)
		So(mapsum.Sum("ap"), ShouldEqual, 3)
		mapsum.Insert("app", 2)
		So(mapsum.Sum("ap"), ShouldEqual, 5)

		// 输入: insert("apple", 3), 输出: Null
		// 输入: sum("ap"), 输出: 3
		// 输入: insert("app", 2), 输出: Null
		// 输入: sum("ap"), 输出: 5

		mapsum = &MapSum{
			Children: map[byte]*MapSum{},
			Map:      make(map[string]int),
		}

		mapsum.Insert("aa", 3)
		So(mapsum.Sum("a"), ShouldEqual, 3)
		mapsum.Insert("aa", 2)
		So(mapsum.Sum("a"), ShouldEqual, 2)
	})
}
