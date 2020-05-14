package leetcode

/*
 * @lc app=leetcode.cn id=622 lang=golang
 *
 * [622] 设计循环队列
 */

// @lc code=start
type MyCircularQueue struct {
	head   int
	tail   int
	length int
	data   []int
}

const null = -1

/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		head:   null,
		tail:   null,
		length: k,
		data:   make([]int, k),
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.head == null {
		this.head = 0
		this.tail = 0
		this.data[0] = value
		return true
	}

	l := this.Len()

	if l == this.length {
		return false
	}

	if this.tail == this.length-1 {
		this.tail = 0
	} else {
		this.tail++
	}
	this.data[this.tail] = value

	return true
}

// 0 1 2 3 4
// t     h

/** Delete an element from the circular queue. Return true if the operation is successful. */
func (this *MyCircularQueue) DeQueue() bool {
	if this.head == null {
		return false
	}

	l := this.Len()

	if l == 1 {
		this.head = null
		this.tail = null
		return true
	}

	if this.head == this.length-1 {
		this.head = 0
	} else {
		this.head++
	}

	return true
}

/** Get the front item from the queue. */
func (this *MyCircularQueue) Front() int {
	if this.head != null {
		return this.data[this.head]
	}
	return -1
}

/** Get the last item from the queue. */
func (this *MyCircularQueue) Rear() int {
	if this.tail != null {
		return this.data[this.tail]
	}
	return -1
}

/** Checks whether the circular queue is empty or not. */
func (this *MyCircularQueue) IsEmpty() bool {
	return this.head == null
}

/** Checks whether the circular queue is full or not. */
func (this *MyCircularQueue) IsFull() bool {
	return this.Len() == this.length
}

func (this *MyCircularQueue) Len() int {
	if this.head == null {
		return 0
	}

	l := this.tail - this.head + 1
	if this.tail < this.head {
		l = this.length - this.head + this.tail + 1
	}

	return l
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
// @lc code=end
