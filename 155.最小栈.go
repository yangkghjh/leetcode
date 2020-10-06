package leetcode

/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
type MinStack struct {
	stack []MinStackNode
}

type MinStackNode struct {
	data int
	min  int
}

/** initialize your data structure here. */
func MinStackConstructor() MinStack {
	m := MinStack{
		stack: []MinStackNode{},
	}

	return m
}

func (this *MinStack) Push(x int) {
	n := MinStackNode{
		data: x,
		min:  x,
	}

	if len(this.stack) > 0 {
		if x > this.stack[len(this.stack)-1].min {
			n.min = this.stack[len(this.stack)-1].min
		}
	}

	this.stack = append(this.stack, n)
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	this.stack = this.stack[0 : len(this.stack)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1].data
}

func (this *MinStack) GetMin() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1].min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
