package offer

// 输入两棵二叉树A和B，判断B是不是A的子结构。(约定空树不是任意一个树的子结构)
// B是A的子结构， 即 A中有出现和B相同的结构和节点值。

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}

	return isSame(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

// 推出条件：
// b 为空，完成匹配
// a 为空，匹配失败
// a.Val != b.Val，匹配失败
func isSame(a, b *TreeNode) bool {
	if b == nil {
		return true
	}

	if a == nil || a.Val != b.Val {
		return false
	}

	return isSame(a.Left, b.Left) && isSame(a.Right, b.Right)
}
