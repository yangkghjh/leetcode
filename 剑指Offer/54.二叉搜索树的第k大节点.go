package offer

// 给定一棵二叉搜索树，请找出其中第k大的节点。

func kthLargest(root *TreeNode, k int) int {
	nums := []int{}
	stack := []*TreeNode{}

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Right
		}

		top := len(stack) - 1
		nums = append(nums, stack[top].Val)
		if len(nums) == k {
			return nums[k-1]
		}
		root = stack[top].Left
		stack = stack[:top]
	}

	return -1
}
