package leetcode

import (
	"bytes"
	"container/heap"
	"sort"
)

/*
 * @lc app=leetcode.cn id=451 lang=golang
 *
 * [451] 根据字符出现频率排序
 */

// @lc code=start
func frequencySort(s string) string {
	hash := map[byte]int{}

	for _, l := range []byte(s) {
		_, ok := hash[l]

		if !ok {
			hash[l] = 0
		}

		hash[l]++
	}

	// 快排
	items := make([]byte, len(hash))
	i := 0
	for l := range hash {
		items[i] = l
		i++
	}

	sort.Slice(items, func(i, j int) bool {
		return hash[items[i]] > hash[items[j]]
	})

	buf := new(bytes.Buffer)

	for _, l := range items {
		for i := 0; i < hash[l]; i++ {
			buf.WriteByte(l)
		}
	}

	return buf.String()
}

// 堆排 L+Nlog(N) N 为 字母数量， L 为字符串长度

func frequencySortHeap(s string) string {
	hash := map[byte]int{}

	for _, l := range []byte(s) {
		_, ok := hash[l]

		if !ok {
			hash[l] = 0
		}

		hash[l]++
	}

	h := new(frequencySortItemsSet)
	heap.Init(h)

	for l, n := range hash {
		heap.Push(h, &frequencySortItem{
			Letter: l,
			Val:    n,
		})
	}

	buf := new(bytes.Buffer)

	for h.Len() > 0 {
		item := heap.Pop(h).(*frequencySortItem)
		for i := 0; i < item.Val; i++ {
			buf.WriteByte(item.Letter)
		}
	}

	return buf.String()
}

type frequencySortItem struct {
	Letter byte
	Val    int
}

type frequencySortItemsSet []*frequencySortItem

func (s frequencySortItemsSet) Len() int {
	return len(s)
}

func (s frequencySortItemsSet) Less(i, j int) bool {
	return s[i].Val > s[j].Val
}

func (s frequencySortItemsSet) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *frequencySortItemsSet) Push(x interface{}) {
	*s = append(*s, x.(*frequencySortItem))
}

func (s *frequencySortItemsSet) Pop() interface{} {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return x
}

// @lc code=end
