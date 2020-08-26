package offer

import "math"

// 输入一棵二叉树的根节点，判断该树是不是平衡二叉树。如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。

func isBalanced(root *TreeNode) bool {

	return isBalancedHelper(root) != -1
}

func isBalancedHelper(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := isBalancedHelper(root.Left)
	right := isBalancedHelper(root.Right)

	if left != -1 && right != -1 && math.Abs(float64(left-right)) <= 1 {
		if left > right {
			return left + 1
		}
		return right + 1
	}

	return -1
}
