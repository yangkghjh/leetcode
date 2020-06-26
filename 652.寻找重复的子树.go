package leetcode

import "strconv"

/*
 * @lc app=leetcode.cn id=652 lang=golang
 *
 * [652] 寻找重复的子树
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
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	hash := map[string]int{}
	ans := []*TreeNode{}

	findDuplicateSubtreesSerialize(root, hash, &ans)

	return ans
}

func findDuplicateSubtreesSerialize(node *TreeNode, hash map[string]int, ans *[]*TreeNode) string {
	if node == nil {
		return "#"
	}

	serial := strconv.Itoa(node.Val) + "," + findDuplicateSubtreesSerialize(node.Left, hash, ans) + "," + findDuplicateSubtreesSerialize(node.Right, hash, ans)

	s, ok := hash[serial]
	if ok {
		if s == 1 {
			*ans = append(*ans, node)
		}
		hash[serial]++
	} else {
		hash[serial] = 1
	}

	return serial
}

// @lc code=end
