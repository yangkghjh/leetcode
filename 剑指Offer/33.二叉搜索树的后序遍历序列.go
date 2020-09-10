package offer

// 输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。

func verifyPostorder(postorder []int) bool {
	return verifyPostorderHelper(postorder, 0, len(postorder)-1)
}

func verifyPostorderHelper(postorder []int, i, j int) bool {
	if j-i <= 1 {
		return true
	}

	p := i
	for postorder[p] < postorder[j] {
		p++
	}
	m := p
	for postorder[p] > postorder[j] {
		p++
	}

	return p == j && verifyPostorderHelper(postorder, i, m-1) && verifyPostorderHelper(postorder, m, j-1)
}
