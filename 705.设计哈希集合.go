package leetcode

/*
 * @lc app=leetcode.cn id=705 lang=golang
 *
 * [705] 设计哈希集合
 */

// @lc code=start
type MyHashSet struct {
	table []byte
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{
		table: make([]byte, 125001),
	}
}

func (this *MyHashSet) Add(key int) {
	this.table[key/8] |= 0x01 << (key % 8)
}

func (this *MyHashSet) Remove(key int) {
	if this.Contains(key) {
		this.table[key/8] &= 0xFE<<(key%8) | 0xFE>>(8-key%8)
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	x := this.table[key/8]
	return x>>(key%8)&1 == 1
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
// @lc code=end
