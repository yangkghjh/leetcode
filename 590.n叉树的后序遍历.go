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
		stack = stack[:l]

		if node == nil {
			continue
		}

		ans = append(ans, node.Val)

		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}

	return ans
}

// @lc code=end
