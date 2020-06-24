package leetcode

/*
 * @lc app=leetcode.cn id=706 lang=golang
 *
 * [706] 设计哈希映射
 */

// @lc code=start
const HashMapModulo = 101

type HashMapNode struct {
	value int
	key   int
	next  *HashMapNode
}

type MyHashMap struct {
	list [HashMapModulo]*HashMapNode
}

/** Initialize your data structure here. */
func NewMyHashMap() MyHashMap {
	// func Constructor() MyHashMap {
	o := MyHashMap{
		list: [HashMapModulo]*HashMapNode{},
	}

	return o
}

func (this *MyHashMap) genHashKey(key int) int {
	return key % HashMapModulo
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	hashKey := this.genHashKey(key)
	node := this.list[hashKey]
	if node == nil {
		this.list[hashKey] = &HashMapNode{
			key:   key,
			value: value,
		}
	} else {
		var prev *HashMapNode
		for node != nil {
			if node.key == key {
				node.value = value
				return
			}

			prev = node
			node = node.next
		}

		prev.next = &HashMapNode{
			key:   key,
			value: value,
		}
	}
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	node := this.list[this.genHashKey(key)]

	for node != nil {
		if node.key == key {
			return node.value
		}

		node = node.next
	}

	return -1
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	hashKey := this.genHashKey(key)
	node := this.list[hashKey]

	if node == nil {
		return
	}

	if node.key == key {
		this.list[hashKey] = node.next
	}

	prev := node
	curr := node.next
	for curr != nil {
		if curr.key == key {
			prev.next = curr.next
			return
		}

		prev = curr
		curr = curr.next
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
// @lc code=end
