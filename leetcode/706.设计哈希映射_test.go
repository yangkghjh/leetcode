package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMyHashMap(t *testing.T) {
	Convey("[706] 设计哈希映射", t, func() {
		hashMap := NewMyHashMap()
		hashMap.Put(1, 1)
		hashMap.Put(2, 2)
		So(hashMap.Get(1), ShouldEqual, 1)  // 返回 1
		So(hashMap.Get(3), ShouldEqual, -1) // 返回 -1 (未找到)
		hashMap.Put(2, 1)                   // 更新已有的值
		So(hashMap.Get(2), ShouldEqual, 1)  // 返回 1
		hashMap.Remove(2)                   // 删除键为2的数据
		So(hashMap.Get(2), ShouldEqual, -1) // 返回 -1 (未找到)
		hashMap.Put(0, 5)
		hashMap.Put(HashMapModulo, 6)
		So(hashMap.Get(0), ShouldEqual, 5)
		So(hashMap.Get(HashMapModulo), ShouldEqual, 6)
	})
}
