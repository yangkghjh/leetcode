package other

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLRUCache(t *testing.T) {
	Convey("LRU 缓存", t, func() {
		cache := &LRUCache{
			hash:     map[int]*LRUNode{},
			capacity: 2,
		}
		cache.Put(1, 1)
		cache.Put(2, 2)
		So(cache.Get(1), ShouldEqual, 1)
		cache.Put(3, 3)
		So(cache.Get(2), ShouldEqual, -1)
		cache.Put(4, 4)
		So(cache.Get(1), ShouldEqual, -1) // 返回 -1 (未找到)
		So(cache.Get(3), ShouldEqual, 3)  // 返回  3
		So(cache.Get(4), ShouldEqual, 4)  // 返回  4
	})
}

func printLRUList(head *LRUNode) {
	if head != nil {
		fmt.Print(head.value, ",")
		printLRUList(head.next)
	} else {
		fmt.Print("\n")
	}
}
