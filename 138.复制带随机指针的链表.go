package leetcode

/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] 复制带随机指针的链表
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *CopyRandomListNode) *CopyRandomListNode {
	nodePair := map[*CopyRandomListNode]*CopyRandomListNode{}
	ans := new(CopyRandomListNode)
	cur := ans
	for p := head; p != nil; p = p.Next {
		q, ok := nodePair[p]
		if !ok {
			q = &CopyRandomListNode{
				Val: p.Val,
			}

			nodePair[p] = q
		}
		cur.Next = q
		cur = cur.Next

		if p.Random != nil {
			r, ok := nodePair[p.Random]
			if !ok {
				r = &CopyRandomListNode{
					Val: p.Random.Val,
				}

				nodePair[p.Random] = r
			}
			q.Random = r
		}
	}

	return ans.Next
}

// 时间复杂度： O(N)
// 空间复杂度： O(N)

// @lc code=end

// Definition for a Node.
type CopyRandomListNode struct {
	Val    int
	Next   *CopyRandomListNode
	Random *CopyRandomListNode
}
