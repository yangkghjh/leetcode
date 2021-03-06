package leetcode

/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for {
		if fast == nil || fast.Next == nil {
			return false
		}

		if fast == slow {
			return true
		}

		fast = fast.Next.Next
		slow = slow.Next
	}
}

// @lc code=end
