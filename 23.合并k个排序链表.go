package main

/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeLists(lists, 0, len(lists)-1)
}

func mergeLists(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}

	mid := (left + right) / 2
	list1 := mergeLists(lists, left, mid)
	list2 := mergeLists(lists, mid+1, right)
	return _mergeTwoLists(list1, list2)
}

func _mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
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
