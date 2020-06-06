package leetcode

/*
 * @lc app=leetcode.cn id=707 lang=golang
 *
 * [707] 设计链表
 */

// @lc code=start
type MyLinkedList struct {
	Head *LinkedListNode
}

type LinkedListNode struct {
	Val  int
	Next *LinkedListNode
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if this.Head == nil {
		return -1
	}
	i := 0
	p := this.Head
	for {
		if i == index {
			return p.Val
		}

		if p.Next == nil {
			return -1
		}
		p = p.Next
		i++
	}
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	this.Head = &LinkedListNode{
		Val:  val,
		Next: this.Head,
	}
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	p := this.Head
	if p == nil {
		this.Head = &LinkedListNode{
			Val: val,
		}
		return
	}
	for {
		if p.Next == nil {
			p.Next = &LinkedListNode{
				Val: val,
			}
			return
		}

		p = p.Next
	}
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	}

	if this.Head == nil {
		return
	}

	i := 1
	p1 := this.Head
	p2 := this.Head.Next

	for {
		if i == index {
			p1.Next = &LinkedListNode{
				Val:  val,
				Next: p2,
			}
		}

		if p2 == nil {
			return
		}

		p1 = p2
		p2 = p2.Next
		i++
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index == 0 {
		this.Head = this.Head.Next
		return
	}

	i := 1
	p := this.Head
	for {
		if i == index {
			if p.Next != nil {
				p.Next = p.Next.Next
			}

			return
		}

		if p.Next == nil {
			return
		}

		p = p.Next
		i++
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
// @lc code=end
