package leetcode

/*
 * @lc app=leetcode.cn id=211 lang=golang
 *
 * [211] 添加与搜索单词 - 数据结构设计
 */

// @lc code=start
type WordDictionary struct {
	Children map[byte]*WordDictionary
	IsEnd    bool
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		Children: make(map[byte]*WordDictionary),
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	node := this

	for _, l := range []byte(word) {
		n, ok := node.Children[l]

		if !ok {
			n = &WordDictionary{
				Children: make(map[byte]*WordDictionary),
			}
			node.Children[l] = n
		}

		node = n
	}

	node.IsEnd = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	nodes := []*WordDictionary{this}
	next := []*WordDictionary{}

	for _, l := range []byte(word) {
		if l == '.' {
			for _, node := range nodes {
				for _, n := range node.Children {
					next = append(next, n)
				}
			}
		} else {
			for _, node := range nodes {
				n, ok := node.Children[l]
				if ok {
					next = append(next, n)
				}
			}
		}

		nodes = next
		next = []*WordDictionary{}
	}

	for _, n := range nodes {
		if n.IsEnd {
			return true
		}
	}

	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end
