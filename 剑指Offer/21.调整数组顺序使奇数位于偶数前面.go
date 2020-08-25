package offer

// 输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。

func exchange(nums []int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		if nums[left]%2 == 0 {
			nums[left], nums[right] = nums[right], nums[left]
			right--
		} else {
			left++
		}
	}

	return nums
}
