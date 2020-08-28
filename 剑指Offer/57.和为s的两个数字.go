package offer

// 输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。如果有多对数字的和等于s，则输出任意一对即可。

func twoSumInSlice(nums []int, target int) []int {
	l := len(nums)
	left, right := 0, l-1

	for left < right && left >= 0 && right < l {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{nums[left], nums[right]}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}

	return []int{}
}
