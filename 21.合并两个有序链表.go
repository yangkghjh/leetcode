package main

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	pointer := &ListNode{}
	pre := &ListNode{Next: pointer}
	for {
		if l1 == nil {
			pointer.Next = l2
			break
		}
		if l2 == nil {
			pointer.Next = l1
			break
		}
		if l1.Val <= l2.Val {
			pointer.Next = l1
			pointer = pointer.Next
			l1 = l1.Next
		} else {
			pointer.Next = l2
			pointer = pointer.Next
			l2 = l2.Next
		}
	}

	return pre.Next.Next
}

// @lc code=end
