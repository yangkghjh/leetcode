package offer

// 27.二叉树的镜像
// 请完成一个函数，输入一个二叉树，该函数输出它的镜像。

func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := root.Left
	root.Left = mirrorTree(root.Right)
	root.Right = mirrorTree(left)

	return root
}
