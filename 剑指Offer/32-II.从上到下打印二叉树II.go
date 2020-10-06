package offer

// 从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。

func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	stack := []*TreeNode{root}
	ans := [][]int{}

	for len(stack) > 0 {
		next := []*TreeNode{}
		vals := []int{}

		for _, node := range stack {
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
	}

	return ans
}
