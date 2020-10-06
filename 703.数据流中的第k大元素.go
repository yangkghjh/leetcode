package leetcode

import "container/heap"

/*
 * @lc app=leetcode.cn id=703 lang=golang
 *
 * [703] 数据流中的第K大元素
 */

// @lc code=start
type KthLargest struct {
	nums *SmallTopHeap
	len  int
}

type SmallTopHeap []int

func (h *SmallTopHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *SmallTopHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h SmallTopHeap) Len() int {
	return len(h)
}

func (h SmallTopHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h SmallTopHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// func Constructor(k int, nums []int) KthLargest {
func NewKthLargest(k int, nums []int) KthLargest {
	h := &KthLargest{
		len:  k,
		nums: new(SmallTopHeap),
	}
	for _, n := range nums {
		h.Add(n)
	}

	return *h
}

func (this *KthLargest) Add(val int) int {
	if this.nums.Len() < this.len {
		heap.Push(this.nums, val)
	} else if (*this.nums)[0] < val {
		heap.Pop(this.nums)
		heap.Push(this.nums, val)
	}

	if this.nums.Len() == 0 {
		return -1
	}

	return (*this.nums)[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
// @lc code=end
