package offer

import "container/list"

// 请定义一个队列并实现函数 max_value 得到队列里的最大值，要求函数max_value、push_back 和 pop_front 的均摊时间复杂度都是O(1)。

// 若队列为空，pop_front 和 max_value 需要返回 -1

// MaxQueue 最大队列
type MaxQueue struct {
	List *list.List
	Nums []int
}

// func Constructor() MaxQueue {
func NewMaxQueue() MaxQueue {
	return MaxQueue{
		List: list.New(),
		Nums: []int{},
	}
}

func (this *MaxQueue) Max_value() int {
	if this.List.Len() == 0 {
		return -1
	}

	return this.List.Front().Value.(int)
}

func (this *MaxQueue) Push_back(value int) {
	for this.List.Len() > 0 && value > this.List.Back().Value.(int) {
		this.List.Remove(this.List.Back())
	}

	this.List.PushBack(value)
	this.Nums = append(this.Nums, value)
}

func (this *MaxQueue) Pop_front() int {
	if len(this.Nums) == 0 {
		return -1
	}

	n := this.Nums[0]
	this.Nums = this.Nums[1:]

	if n == this.List.Front().Value.(int) {
		this.List.Remove(this.List.Front())
	}

	return n
}
