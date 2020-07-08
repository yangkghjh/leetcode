package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=110 lang=golang
 *
 * [110] 平衡二叉树
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
func isBalanced(root *TreeNode) bool {
	return isBalancedRecur(root) != -1
}

func isBalancedRecur(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := isBalancedRecur(node.Left)
	if left == -1 {
		return -1
	}

	right := isBalancedRecur(node.Right)
	if right == -1 {
		return -1
	}

	if math.Abs(float64(left-right)) > 1 {
		return -1
	}

	if left > right {
		return left + 1
	}

	return right + 1
}

// @lc code=end
