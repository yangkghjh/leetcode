package leetcode

/*
 * @lc app=leetcode.cn id=677 lang=golang
 *
 * [677] 键值映射
 */

// @lc code=start
type MapSum struct {
	Children map[byte]*MapSum
	Num      int
	Map      map[string]int
}

/** Initialize your data structure here. */
func NewMapSum() MapSum {
	return MapSum{
		Children: make(map[byte]*MapSum),
		Map:      make(map[string]int),
	}
}

func (this *MapSum) Insert(key string, val int) {
	if v, ok := this.Map[key]; ok {
		val -= v
	}
	this.Map[key] = val

	node := this

	for _, l := range []byte(key) {
		n, ok := node.Children[l]

		if !ok {
			n = &MapSum{
				Children: make(map[byte]*MapSum),
			}
			node.Children[l] = n
		}

		n.Num += val
		node = n
	}
}

func (this *MapSum) Sum(prefix string) int {
	node := this

	for _, l := range []byte(prefix) {
		n, ok := node.Children[l]

		if !ok {
			return 0
		}

		node = n
	}

	return node.Num
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
// @lc code=end
