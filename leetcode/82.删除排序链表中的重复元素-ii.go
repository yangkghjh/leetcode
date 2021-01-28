package leetcode

/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	prev := &ListNode{
		Next: head,
	}
	ans := prev
	dup := false

	// [0] 1 1 2 3
	// [0] 1 1 1 2 3
	for head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			dup = true
			head.Next = head.Next.Next
			continue
		}

		if dup {
			head = head.Next
			prev.Next = head
			dup = false
			continue
		}
		prev = head
		head = head.Next
	}

	if dup {
		prev.Next = nil
	}

	return ans.Next
}

// @lc code=end
