package leetcode

// ListNode 链表
type ListNode struct {
	Val  int
	Next *ListNode
}

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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

func NewBinaryTree(values ...int) *TreeNode {
	if len(values) == 0 {
		return nil
	}

	root := &TreeNode{Val: values[0]}
	stack := []*TreeNode{root}

	dep := 1
	i := 1
	var left, right *TreeNode

	for k := 0; k < dep; {
		j := k / 2
		if k%2 == 1 {
			j++
		}
		// fmt.Println(dep, k, j)

		if stack[j] != nil {

			left = newTreeNode(values[i])
			if i+1 < len(values) {
				right = newTreeNode(values[i+1])
			} else {
				right = nil
			}

			stack[j].Left = left
			stack[j].Right = right
			stack[j] = left
			stack = append(stack, right)

			i += 2
			if i >= len(values) {
				break
			}
		} else {
			stack[j] = nil
			stack = append(stack, nil)
		}

		if k+1 == dep {
			k = 0
			dep *= 2
		} else {
			k++
		}
	}

	return root
}

func newTreeNode(value int) *TreeNode {
	if value != -1 {
		return &TreeNode{
			Val: value,
		}
	}

	return nil
}
