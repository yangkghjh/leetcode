package offer

import "container/heap"

// 输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

func getLeastNumbers(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}
	h := new(Heap)

	for _, num := range arr {
		if h.Len() < k {
			heap.Push(h, num)
		} else {
			if num < (*h)[0] {
				heap.Pop(h)
				heap.Push(h, num)
			}
		}
	}

	return *h
}

// Heap 堆
type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] > h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push 推入元素
func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop 弹出元素
func (h *Heap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]

	return x
}
