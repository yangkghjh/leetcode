package leetcode

/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	for k := range inorder {
		if inorder[k] == preorder[0] {
			return &TreeNode{
				Val:   preorder[0],
				Left:  buildTree(preorder[1:k+1], inorder[:k]),
				Right: buildTree(preorder[k+1:], inorder[k+1:]),
			}
		}
	}

	return nil
}

// 时间复杂度： O(N)
// 空间复杂度： O(N)

// @lc code=end
