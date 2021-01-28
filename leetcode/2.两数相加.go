package leetcode

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := &ListNode{}
	cur := ans
	carry := 0
	for l1 != nil || l2 != nil || carry == 1 {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		cur.Next = &ListNode{
			Val: carry % 10,
		}

		cur = cur.Next
		carry /= 10
	}

	return ans.Next
}

// 时间复杂度： O(max(m,n))
// 空间复杂度： O(max(m,n))

// @lc code=end
