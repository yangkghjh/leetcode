package leetcode

/*
 * @lc app=leetcode.cn id=235 lang=golang
 *
 * [235] 二叉搜索树的最近公共祖先
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	ans := root

	if p.Val > q.Val {
		p, q = q, p
	}

	for ans != nil {
		if p.Val <= ans.Val && q.Val >= ans.Val {
			break
		}

		if p.Val < ans.Val {
			ans = ans.Left
		} else {
			ans = ans.Right
		}
	}

	return ans
}

// 时间复杂度： O(D) D 为二叉树高度
// 空间复杂度： O(1)

// @lc code=end
