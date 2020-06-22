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

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }
