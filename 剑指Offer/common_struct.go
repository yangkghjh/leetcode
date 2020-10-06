package offer

import "math"

// ListNode 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

// NewList 创建链表
func NewList(list ...int) *ListNode {
	node := &ListNode{}
	head := node
	for _, n := range list {
		node.Next = &ListNode{
			Val: n,
		}
		node = node.Next
	}

	return head.Next
}

// TreeNode 二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeNode 创建二叉树节点
func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

// 空节点
const null = math.MinInt64

// NewBinaryTree 根据层序遍历创建二叉树
func NewBinaryTree(values ...int) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	root := &TreeNode{Val: values[0]}
	list := []*TreeNode{root}
	ans := []*TreeNode{}

	i := 1

	for len(list) == 0 || i < len(values) {
		for _, l := range list {
			if i >= len(values) {
				break
			}
			if l != nil {
				l.Left = newTreeNode(values[i])
				if i+1 < len(values) {
					l.Right = newTreeNode(values[i+1])
				}

				ans = append(ans, l.Left, l.Right)

				i += 2
			}
		}
		list = ans
		ans = []*TreeNode{}
	}

	return root
}

func newTreeNode(value int) *TreeNode {
	if value != null {
		return &TreeNode{
			Val: value,
		}
	}

	return nil
}

// TreeNodeDeepEqual 二叉树相等
func TreeNodeDeepEqual(m, n *TreeNode) bool {
	if m == n {
		return true
	}

	if m == nil || n == nil || m.Val != n.Val {
		return false
	}

	return TreeNodeDeepEqual(m.Left, n.Left) && TreeNodeDeepEqual(m.Right, n.Right)
}

// SerializeBinaryTree 序列化二叉树
func SerializeBinaryTree(node *TreeNode) []int {
	list := []*TreeNode{node}
	next := []*TreeNode{}

	ans := []int{}

	for len(list) != 0 {
		for _, n := range list {
			if n != nil {
				ans = append(ans, n.Val)
				next = append(next, n.Left, n.Right)
			} else {
				ans = append(ans, -1)
			}
		}

		list = next
		next = []*TreeNode{}
	}

	for i := len(ans) - 1; i >= 0; i-- {
		if ans[i] == -1 {
			ans = ans[:i]
		} else {
			break
		}
	}

	return ans
}
