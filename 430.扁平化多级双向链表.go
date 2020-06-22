package leetcode

/*
 * @lc app=leetcode.cn id=430 lang=golang
 *
 * [430] 扁平化多级双向链表
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *FlattenNode) *FlattenNode {
	for p := root; p != nil; p = p.Next {
		if p.Child != nil {
			next := p.Next
			p.Next, p.Child.Prev = p.Child, p
			q := p.Child
			for q.Next != nil {
				q = q.Next
			}

			q.Next = next

			if next != nil {
				next.Prev = q
			}
			p.Child = nil
		}
	}

	return root
}

// 时间复杂度： O(N)
// 空间复杂度： O(1)

// @lc code=end

// Definition for a Node.
type FlattenNode struct {
	Val   int
	Prev  *FlattenNode
	Next  *FlattenNode
	Child *FlattenNode
}

func NewFlattenList(list ...int) *FlattenNode {
	node := &FlattenNode{}
	head := node
	for _, n := range list {
		node.Next = &FlattenNode{
			Val: n,
		}
		node = node.Next
	}

	return head.Next
}
