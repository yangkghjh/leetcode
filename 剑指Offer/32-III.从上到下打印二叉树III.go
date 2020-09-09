package offer

// 请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。

func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	stack := []*TreeNode{root}
	ans := [][]int{}
	dir := 1

	for len(stack) > 0 {
		next := []*TreeNode{}
		vals := []int{}

		i := 0
		if dir == -1 {
			i = len(stack) - 1
		}
		for ; i < len(stack) && i >= 0; i += dir {
			node := stack[i]
			if node != nil {
				vals = append(vals, node.Val)
				if node.Left != nil {
					next = append(next, node.Left)
				}
				if node.Right != nil {
					next = append(next, node.Right)
				}
			}
		}

		stack = next
		ans = append(ans, vals)
		dir = -dir
	}

	return ans
}
