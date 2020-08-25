package leetcode

/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	queue := []*TreeNode{root}
	ans := []int{}

	for len(queue) > 0 {
		cur := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		for cur != nil {
			ans = append(ans, cur.Val)

			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}

			cur = cur.Left
		}
	}

	return ans
}

// @lc code=end
