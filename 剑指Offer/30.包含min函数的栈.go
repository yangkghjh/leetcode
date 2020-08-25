package offer

// 定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的 min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。

// MinStack 最小栈
type MinStack struct {
	stack []int
	min   []int
}

// func Constructor() MinStack {

// NewMinStack 最小栈构造函数
func NewMinStack() MinStack {
	return MinStack{
		stack: []int{},
		min:   []int{},
	}
}

// Push 推入数据
func (s *MinStack) Push(x int) {
	min := x
	if len(s.stack) > 0 {
		last := s.min[len(s.min)-1]
		if last < min {
			min = last
		}
	}

	s.stack = append(s.stack, x)
	s.min = append(s.min, min)
}

// Pop 弹出数据
func (s *MinStack) Pop() {
	if len(s.stack) > 0 {
		s.stack = s.stack[:len(s.stack)-1]
		s.min = s.min[:len(s.min)-1]
	}
}

// Top 栈顶
func (s *MinStack) Top() int {
	if len(s.stack) == 0 {
		return -1
	}

	return s.stack[len(s.stack)-1]
}

// Min 最小值
func (s *MinStack) Min() int {
	if len(s.min) == 0 {
		return -1
	}

	return s.min[len(s.min)-1]
}
