package leetcode

/*
 * @lc app=leetcode.cn id=421 lang=golang
 *
 * [421] 数组中两个数的最大异或值
 */

// @lc code=start
func findMaximumXOR(nums []int) int {
	trie := new(MaximumXORTrie)

	for _, n := range nums {
		trie.Insert(n)
	}

	max := 0
	for _, n := range nums {
		ans := trie.MaximumXOR(n)
		if ans > max {
			max = ans
		}
	}

	return max
}

// 时间复杂度： O(2*32*N) -> O(N)

type MaximumXORTrie struct {
	Children [2]*MaximumXORTrie
}

// 时间复杂度： O(32) -> O(1)

func (t *MaximumXORTrie) Insert(num int) {
	node := t

	for i := 31; i >= 0; i-- {
		value := 0
		if ((num >> i) & 1) == 1 {
			value = 1
		}

		next := node.Children[value]
		if next == nil {
			next = new(MaximumXORTrie)
			node.Children[value] = next
		}

		node = next
	}
}

// 时间复杂度： O(32) -> O(1)

func (t *MaximumXORTrie) MaximumXOR(num int) int {
	res := 0
	node := t

	for i := 31; i >= 0; i-- {
		value := 1
		if ((num >> i) & 1) == 1 {
			value = 0
		}

		next := node.Children[value]
		res <<= 1
		if next != nil {
			res += 1
		} else {
			if value == 1 {
				value = 0
			} else {
				value = 1
			}
			next = node.Children[value]
		}

		node = next
	}

	return res
}

// @lc code=end
