package offer

// ListNode 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// NewList 创建链表
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
