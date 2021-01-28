package leetcode

/*
 * @lc app=leetcode.cn id=590 lang=golang
 *
 * [590] N叉树的后序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
	stack := []*Node{root}
	ans := []int{}

	for len(stack) > 0 {
		l := len(stack) - 1
		node := stack[l]

		if node == nil {
			stack = stack[:l]
			continue
		}

		if len(node.Children) == 0 {
			ans = append(ans, node.Val)
			stack = stack[:l]
		} else {
			for i := len(node.Children) - 1; i >= 0; i-- {
				stack = append(stack, node.Children[i])
			}
			node.Children = nil
		}
	}

	return ans
}

// 时间复杂度： O(N)
// 空间复杂度： O(N)

// @lc code=end
