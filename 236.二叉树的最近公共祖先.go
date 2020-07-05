package leetcode

/*
 * @lc app=leetcode.cn id=236 lang=golang
 *
 * [236] 二叉树的最近公共祖先
 */

// @lc code=start
/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
// 方法二：递归找到父节点
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}

	return left
}

// 方法一：保存父节点列表并对比
// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	pt := findBinaryTreeTrace(root, p)
// 	qt := findBinaryTreeTrace(root, q)

// 	fmt.Println(pt, qt)
// 	pl, ql := len(pt)-1, len(qt)-1

// 	for ; pl >= 0 && ql >= 0; pl, ql = pl-1, ql-1 {
// 		if pt[pl] != qt[ql] {
// 			return pt[pl+1]
// 		}
// 	}

// 	if pl < 0 {
// 		return pt[0]
// 	}

// 	return qt[0]
// }

func findBinaryTreeTrace(node, target *TreeNode) []*TreeNode {
	if node == nil {
		return nil
	}

	if node == target {
		return []*TreeNode{node}
	}

	if v := findBinaryTreeTrace(node.Left, target); v != nil {
		return append(v, node)
	}
	if v := findBinaryTreeTrace(node.Right, target); v != nil {
		return append(v, node)
	}

	return nil
}

// 时间复杂度： O(N)
// 空间复杂度： O(N)

// @lc code=end
