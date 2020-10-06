package leetcode

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTrie(t *testing.T) {
	Convey("[208] 实现 Trie (前缀树)", t, func() {
		trie := &Trie{
			Children: map[byte]*Trie{},
		}

		trie.Insert("apple")
		So(trie.Search("apple"), ShouldBeTrue)   // 返回 true
		So(trie.Search("app"), ShouldBeFalse)    // 返回 false
		So(trie.StartsWith("app"), ShouldBeTrue) // 返回 true
		trie.Insert("app")
		So(trie.Search("app"), ShouldBeTrue) // 返回 true
	})
}
