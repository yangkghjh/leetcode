package leetcode

import "container/list"

/*
 * @lc app=leetcode.cn id=225 lang=golang
 *
 * [225] 用队列实现栈
 */

// @lc code=start

// 尝试用 list.List
type MyStack struct {
	queue1 list.List
	queue2 list.List
}

/** Initialize your data structure here. */
func NewMyStack() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	if this.queue2.Front() != nil {
		this.queue2.PushBack(x)
		return
	}

	this.queue1.PushBack(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.Empty() {
		return -1
	}

	var src *list.List
	var dst *list.List
	if this.queue1.Front() == nil {
		src = &this.queue2
		dst = &this.queue1
	} else {
		src = &this.queue1
		dst = &this.queue2
	}

	for src.Front().Next() != nil {
		dst.PushBack(src.Front().Value)
		src.Remove(src.Front())
	}

	value := src.Front().Value
	src.Remove(src.Front())

	return value.(int)
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if this.Empty() {
		return -1
	}

	var src *list.List
	var dst *list.List
	if this.queue1.Front() == nil {
		src = &this.queue2
		dst = &this.queue1
	} else {
		src = &this.queue1
		dst = &this.queue2
	}

	for src.Front().Next() != nil {
		dst.PushBack(src.Front().Value)
		src.Remove(src.Front())
	}

	value := src.Front().Value
	dst.PushBack(src.Front().Value)
	src.Remove(src.Front())

	return value.(int)
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.queue1.Front() == nil && this.queue2.Front() == nil
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
// @lc code=end
