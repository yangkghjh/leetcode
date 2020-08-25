package leetcode

/*
 * @lc app=leetcode.cn id=346 lang=golang
 *
 * [346] 数据流中的移动平均值
 */

// @lc code=start

/// 方法二：循环队列
type MovingAverage struct {
	Queue []int
	Tail  int
	Sum   int
	Size  int
	Len   int
}

/** Initialize your data structure here. */
func MovingAverageConstructor(size int) MovingAverage {
	return MovingAverage{
		Size:  size,
		Queue: make([]int, size),
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if this.Size == 0 {
		return 0
	}

	this.Sum += val
	this.Sum -= this.Queue[this.Tail]
	this.Queue[this.Tail] = val
	this.Tail++
	if this.Tail >= this.Size {
		this.Tail = 0
	}
	if this.Len < this.Size {
		this.Len++
	}

	return float64(this.Sum) / float64(this.Len)
}

// 方法一：双端链表
// type MovingAverage struct {
// 	Head *ListNode
// 	Tail *ListNode
// 	Len  int
// 	Sum  int
// 	Size int
// }

/** Initialize your data structure here. */
// func MovingAverageConstructor(size int) MovingAverage {
// 	return MovingAverage{
// 		Size: size,
// 	}
// }

// func (this *MovingAverage) Next(val int) float64 {
// 	if this.Size == 0 {
// 		return 0
// 	}

// 	n := &ListNode{
// 		Val: val,
// 	}

// 	if this.Len == 0 {
// 		this.Head, this.Tail = n, n
// 		this.Len = 1
// 		this.Sum = val
// 	} else if this.Len < this.Size {
// 		this.Tail.Next = n
// 		this.Tail = n
// 		this.Len++
// 		this.Sum += val
// 	} else {
// 		this.Sum -= this.Head.Val
// 		this.Head = this.Head.Next
// 		this.Tail.Next = n
// 		this.Tail = n
// 		this.Sum += val
// 	}

// 	return float64(this.Sum) / float64(this.Len)
// }

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
// @lc code=end
