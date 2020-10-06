package leetcode

/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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

// 迭代
// func postorderTraversal(root *TreeNode) []int {
// 	ans := []int{}
// 	stack := []*TreeNode{root}

// 	for len(stack) > 0 {
// 		node := stack[len(stack)-1]
// 		stack = stack[:len(stack)-1]

// 		if node == nil {
// 			continue
// 		}

// 		ans = append([]int{node.Val}, ans...)

// 		stack = append(stack, node.Left, node.Right)
// 	}

// 	return ans
// }

// 时间复杂度： O(N)
// 空间复杂度： O(N)

// 递归

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	ans := []int{}

	ans = append(ans, postorderTraversal(root.Left)...)
	ans = append(ans, postorderTraversal(root.Right)...)

	ans = append(ans, root.Val)

	return ans
}

// 时间复杂度： O(N)
// 空间复杂度： O(N)

// @lc code=end
