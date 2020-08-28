package offer

import "container/list"

// 给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

func maxSlidingWindow(nums []int, k int) []int {
	l := len(nums)
	if l == 0 {
		return []int{}
	}
	ans := []int{}
	moq := []int{}

	if k > l {
		k = l
	}

	for i := 0; i < l; i++ {
		n := nums[i]
		for len(moq) > 0 && n > moq[len(moq)-1] {
			moq = moq[:len(moq)-1]
		}
		moq = append(moq, n)

		if i >= k-1 {
			ans = append(ans, moq[0])
			m := nums[i-k+1]
			if m == moq[0] {
				moq = moq[1:]
			}
		}
	}

	return ans
}

// 单调队列
func maxSlidingWindowMoQueue(nums []int, k int) []int {
	l := len(nums)
	if l == 0 {
		return []int{}
	}
	ans := []int{}
	moq := NewMoQueue()

	if k > l {
		k = l
	}

	for i := 0; i < k-1; i++ {
		moq.Push(nums[i])
	}

	for i := k - 1; i < l; i++ {
		moq.Push(nums[i])
		ans = append(ans, moq.Max())
		moq.Pop(nums[i-k+1])
	}

	return ans
}

// MoQueue 单调队列，双端队列实现
type MoQueue struct {
	*list.List // 双端队列
	// 会用到的函数
	// PushFront(x interface{})
	// PushBack(x interface{})
	// Front() *Element
	// Back() *Element
	// Remove(el *Element)
	// Len()
}

// NewMoQueue 初始化
func NewMoQueue() *MoQueue {
	return &MoQueue{list.New()}
}

// Push 推入一个元素
func (m *MoQueue) Push(n int) {
	for m.Len() > 0 &&
		n > m.Back().Value.(int) {
		m.Remove(m.Back())
	}
	m.PushBack(n)
}

// Max 最大值
func (m *MoQueue) Max() int {
	return m.Front().Value.(int)
}

// Pop 弹出一个元素
func (m *MoQueue) Pop(n int) {
	if n == m.Front().Value.(int) {
		m.Remove(m.Front())
	}
}
