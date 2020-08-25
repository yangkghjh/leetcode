package leetcode

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var i int
	pre := &ListNode{
		Next: head,
	} // 避免去除一个节点时无法操作
	fast := pre // 快指针
	slow := pre // 慢指针
	for {
		if fast.Next != nil {
			fast = fast.Next
		} else {
			if i == 0 {
				return nil
			}
			slow.Next = slow.Next.Next
			return pre.Next
		}
		if i >= n {
			slow = slow.Next
		}
		i++
	}
}

// @lc code=end
