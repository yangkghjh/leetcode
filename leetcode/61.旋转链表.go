package leetcode

/*
 * @lc app=leetcode.cn id=61 lang=golang
 *
 * [61] 旋转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}

	l := 0
	for p := head; p != nil; p = p.Next {
		l++
	}

	k %= l

	if k == 0 {
		return head
	}

	p, q := head, head
	for ; k > 0; k-- {
		p = p.Next
	}
	for ; p.Next != nil; p = p.Next {
		q = q.Next
	}

	ans := q.Next
	q.Next = nil
	p.Next = head

	return ans
}

// @lc code=end
