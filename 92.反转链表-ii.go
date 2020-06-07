package leetcode

/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}
	if m >= n || m < 0 {
		return head
	}

	head = &ListNode{
		Next: head,
	}
	start := head
	prev := &ListNode{
		Next: head,
	}
	for i := 0; ; i++ {
		if i == m {
			break
		}
		start = start.Next
		prev = prev.Next
	}

	prev.Next = reverseByLength(start, n-m)
	return head.Next
}

func reverseByLength(head *ListNode, n int) *ListNode {
	res := &ListNode{
		Next: head,
	}

	for i := 0; i < n; i++ {
		if head.Next == nil {
			break
		}

		x := res.Next
		res.Next = head.Next
		head.Next = res.Next.Next
		res.Next.Next = x
	}

	return res.Next
}

// @lc code=end
