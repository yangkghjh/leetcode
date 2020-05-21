package leetcode

/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 */

// @lc code=start
type MyQueue struct {
	push []int
	pop  []int
}

/** Initialize your data structure here. */
func NewMyQueue() MyQueue {
	return MyQueue{
		push: []int{},
		pop:  []int{},
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	if len(this.pop) > 0 {
		this.shift()
	}

	this.push = append(this.push, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.Empty() {
		return -1
	}
	if len(this.push) > 0 {
		this.shift()
	}

	x := this.pop[len(this.pop)-1]
	this.pop = this.pop[:len(this.pop)-1]

	return x
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.Empty() {
		return -1
	}
	if len(this.push) > 0 {
		this.shift()
	}

	x := this.pop[len(this.pop)-1]

	return x
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.push) == 0 && len(this.pop) == 0
}

func (this *MyQueue) shift() {
	if this.Empty() {
		return
	}

	if len(this.push) > 0 {
		for i := len(this.push) - 1; i >= 0; i-- {
			this.pop = append(this.pop, this.push[i])
		}
		this.push = []int{}
		return
	}

	for i := len(this.pop) - 1; i >= 0; i-- {
		this.push = append(this.push, this.pop[i])
	}
	this.pop = []int{}
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end
