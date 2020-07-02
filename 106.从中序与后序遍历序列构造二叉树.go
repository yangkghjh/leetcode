package leetcode

/*
 * @lc app=leetcode.cn id=106 lang=golang
 *
 * [106] 从中序与后序遍历序列构造二叉树
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	l := len(postorder) - 1
	if l < 0 {
		return nil
	}

	for k := range inorder {
		if inorder[k] == postorder[l] {
			return &TreeNode{
				Val:   postorder[l],
				Left:  buildTree(inorder[:k], postorder[:k]),
				Right: buildTree(inorder[k+1:], postorder[k:l]),
			}
		}
	}

	return nil
}

// 时间复杂度： O(N)
// 空间复杂度： O(D)

// @lc code=end
