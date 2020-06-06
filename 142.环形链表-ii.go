package leetcode

/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	fast := head
	slow := head

	for {
		if fast == nil || fast.Next == nil {
			return nil
		}

		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			break
		}
	}

	res := head

	for {
		if res == fast {
			return res
		}

		res = res.Next
		fast = fast.Next
	}
}

// @lc code=end
