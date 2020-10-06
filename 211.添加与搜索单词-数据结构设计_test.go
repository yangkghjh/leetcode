package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWordDictionary(t *testing.T) {
	Convey("[211] 添加与搜索单词 - 数据结构设计", t, func() {
		trie := &WordDictionary{
			Children: make(map[byte]*WordDictionary),
		}

		trie.AddWord("bad")
		trie.AddWord("dad")
		trie.AddWord("mad")
		So(trie.Search("pad"), ShouldBeFalse)
		So(trie.Search("bad"), ShouldBeTrue)
		So(trie.Search(".ad"), ShouldBeTrue)
		So(trie.Search("b.."), ShouldBeTrue)
	})
}
