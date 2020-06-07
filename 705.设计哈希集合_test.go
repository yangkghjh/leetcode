package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMyHashSet(t *testing.T) {
	Convey("[705] 设计哈希集合", t, func() {
		hashSet := Constructor()
		hashSet.Add(1)
		hashSet.Add(2)
		So(hashSet.Contains(1), ShouldBeTrue)  // 返回 true
		So(hashSet.Contains(3), ShouldBeFalse) // 返回 false (未找到)
		hashSet.Add(2)
		So(hashSet.Contains(2), ShouldBeTrue) // 返回 true
		hashSet.Remove(2)
		So(hashSet.Contains(2), ShouldBeFalse) // 返回  false (已经被删除)
		hashSet.Add(1000000)
		hashSet.Add(0)
		So(hashSet.Contains(1000000), ShouldBeTrue) // 返回 true
		So(hashSet.Contains(0), ShouldBeTrue)       // 返回 true
	})
}
