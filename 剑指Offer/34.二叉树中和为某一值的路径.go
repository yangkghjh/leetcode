package offer

// 输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。

func pathSum(root *TreeNode, sum int) [][]int {
	result := [][]int{}
	pathSumHelper(root, sum, []int{}, &result)
	return result
}

func pathSumHelper(root *TreeNode, sum int, trace []int, result *[][]int) {
	if root == nil {
		return
	}

	v := sum - root.Val
	trace = append(trace, root.Val)

	if root.Left == nil && root.Right == nil {
		if v == 0 {
			*result = append(*result, append([]int{}, trace...))
		}
		return
	}

	if root.Left != nil {
		pathSumHelper(root.Left, v, trace, result)
	}

	if root.Right != nil {
		pathSumHelper(root.Right, v, trace, result)
	}
}

func pathSum1(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		if sum-root.Val == 0 {
			return [][]int{[]int{sum}}
		}
		return nil
	}

	ans := [][]int{}

	if root.Left != nil {
		l := pathSum(root.Left, sum-root.Val)
		if l != nil {
			for _, i := range l {
				ans = append(ans, append([]int{root.Val}, i...))
			}
		}
	}

	if root.Right != nil {
		l := pathSum(root.Right, sum-root.Val)
		if l != nil {
			for _, i := range l {
				ans = append(ans, append([]int{root.Val}, i...))
			}
		}
	}

	return ans
}
