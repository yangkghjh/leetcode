package leetcode

/*
 * @lc app=leetcode.cn id=450 lang=golang
 *
 * [450] 删除二叉搜索树中的节点
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
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		if root.Left == nil {
			return root.Right
		}

		if root.Right == nil {
			return root.Left
		}

		min := root.Right
		for min.Left != nil {
			min = min.Left
		}
		min.Left = root.Left
		root = root.Right
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		root.Right = deleteNode(root.Right, key)
	}

	return root
}

// 时间复杂度： O(D)
// 空间复杂度： O(D)

// @lc code=end
