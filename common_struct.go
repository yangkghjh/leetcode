package leetcode

// ListNode 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(list ...int) *ListNode {
	node := &ListNode{}
	head := node
	for _, n := range list {
		node.Next = &ListNode{
			Val: n,
		}
		node = node.Next
	}

	return head.Next
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

func NewBinaryTree(values ...int) *TreeNode {
	root := &TreeNode{Val: values[0]}
	values = values[1:]

	var Q []*TreeNode
	Q = append(Q, root)
	for len(Q) > 0 && len(values) > 0 {
		// pop  tree node
		n := 1
		p := Q[0]
		Q = Q[1:]

		if values[0] != -1 {
			p.Left = &TreeNode{Val: values[0]}
			Q = append(Q, p.Left)
		}
		// 为空的情况可以省略，模式指针就是nil

		// 第二个元素为右子树
		if len(values) >= 2 && values[1] != -1 {
			p.Right = &TreeNode{Val: values[1]}
			Q = append(Q, p.Right)
			n++
		}

		// pop出两个元素
		values = values[n:]
	}
	return root
}
