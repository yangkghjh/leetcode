package leetcode

/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	stack := []*Node{root}
	next := []*Node{}
	ans := [][]int{}

	for len(stack) > 0 {
		cur := []int{}
		for _, node := range stack {
			if node == nil {
				continue
			}

			cur = append(cur, node.Val)

			for _, child := range node.Children {
				next = append(next, child)
			}
		}

		ans = append(ans, cur)
		stack = next
		next = []*Node{}
	}

	return ans
}

// 时间复杂度： O(N)
// 空间复杂度： O(N)

// @lc code=end
