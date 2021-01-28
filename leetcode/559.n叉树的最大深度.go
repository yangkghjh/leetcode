package leetcode

/*
 * @lc app=leetcode.cn id=559 lang=golang
 *
 * [559] N叉树的最大深度
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	max := 0

	for _, child := range root.Children {
		d := maxDepth(child)
		if d > max {
			max = d
		}
	}

	return max + 1
}

// 时间复杂度： O(N)
// 空间复杂度： O(N)

// @lc code=end
