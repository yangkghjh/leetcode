package leetcode

/*
 * @lc app=leetcode.cn id=701 lang=golang
 *
 * [701] 二叉搜索树中的插入操作
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
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: 5}
	}

	node := root
	for {
		if node.Val == val {
			break
		} else if val < node.Val {
			if node.Left != nil {
				node = node.Left
			} else {
				node.Left = &TreeNode{Val: val}
				break
			}
		} else {
			if node.Right != nil {
				node = node.Right
			} else {
				node.Right = &TreeNode{Val: val}
				break
			}
		}
	}

	return root
}

// 时间复杂度： O(D) D 为二叉树的高度
// 空间复杂度： O(1)

// @lc code=end
