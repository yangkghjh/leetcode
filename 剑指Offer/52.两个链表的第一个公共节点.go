package offer

// 输入两个链表，找出它们的第一个公共节点。

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB

	for a != b {
		if a == nil {
			a = headB
		} else {
			a = a.Next
		}

		if b == nil {
			b = headA
		} else {
			b = b.Next
		}
	}

	return a
}
