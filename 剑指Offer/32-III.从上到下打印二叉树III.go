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

		for i := len(stack) - 1; i >= 0; i-- {
			node := stack[i]
			if node != nil {
				vals = append(vals, node.Val)
				if dir == 1 {
					if node.Left != nil {
						next = append(next, node.Left)
					}
					if node.Right != nil {
						next = append(next, node.Right)
					}
				} else {
					if node.Right != nil {
						next = append(next, node.Right)
					}
					if node.Left != nil {
						next = append(next, node.Left)
					}
				}
			}
		}

		stack = next
		ans = append(ans, vals)
		dir = -dir
	}

	return ans
}

// next := []*TreeNode{}
// vals := []int{}

// for i := 0; i < len(stack); i++ {
// 	node := stack[i]
// 	if node != nil {
// 		if dir == 1 {
// 			vals = append(vals, node.Val)
// 		} else {
// 			vals = append([]int{node.Val}, vals...)
// 		}
// 		if node.Left != nil {
// 			next = append(next, node.Left)
// 		}
// 		if node.Right != nil {
// 			next = append(next, node.Right)
// 		}
// 	}
// }

// stack = next
// ans = append(ans, vals)
// dir = -dir
