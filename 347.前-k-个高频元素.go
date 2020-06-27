package leetcode

import "container/heap"

/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	hash := map[int]int{}

	for _, n := range nums {
		_, ok := hash[n]
		if ok {
			hash[n]++
		} else {
			hash[n] = 1
		}
	}

	h := &NodeHeap{}
	if len(hash) < k {
		k = len(hash)
	}
	for v, n := range hash {
		if h.Len() < k {
			heap.Push(h, &MinHeapNode{val: v, times: n})
		} else {
			if n > (*h)[0].times {
				heap.Pop(h)
				heap.Push(h, &MinHeapNode{val: v, times: n})
			}
		}
	}

	ans := []int{}
	for h.Len() > 0 {
		ans = append(ans, h.Pop().(*MinHeapNode).val)
	}

	return ans
}

type MinHeapNode struct {
	val   int
	times int
}

type NodeHeap []*MinHeapNode

func (h NodeHeap) Len() int {
	return len(h)
}

func (h NodeHeap) Less(i, j int) bool {
	return h[i].times < h[j].times
}

func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*MinHeapNode))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 时间复杂度： O(N+NlogK)
// 空间复杂度： O(N)

// @lc code=end
