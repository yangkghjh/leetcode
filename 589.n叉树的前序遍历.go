package leetcode

/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
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

// 时间复杂度： O(N) N 为节点数量
// 空间复杂度： O(N)

// @lc code=end

type Node struct {
	Val      int
	Children []*Node
}
