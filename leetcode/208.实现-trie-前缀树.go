package leetcode

/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
type Trie struct {
	Children map[byte]*Trie
	Finish   bool
}

/** Initialize your data structure here. */
func NewTrie() Trie {
	return Trie{
		Children: map[byte]*Trie{},
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this
	for _, l := range []byte(word) {
		n, ok := node.Children[l]

		if !ok {
			n = &Trie{
				Children: map[byte]*Trie{},
			}
			node.Children[l] = n
		}

		node = n
	}

	node.Finish = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this
	for _, l := range []byte(word) {
		n, ok := node.Children[l]

		if !ok {
			return false
		}

		node = n
	}

	return node.Finish
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, l := range []byte(prefix) {
		n, ok := node.Children[l]

		if !ok {
			return false
		}

		node = n
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
