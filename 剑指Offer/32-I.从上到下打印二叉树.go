package offer

// 从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{root}
	ans := []int{}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		ans = append(ans, node.Val)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return ans
}
