package offer

// 统计一个数字在排序数组中出现的次数。

func search(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	i := 0
	j := l - 1
	for i <= j {
		mid := j - (j-i)/2
		if nums[mid] <= target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	right := i

	i = 0
	j = l - 1
	for i <= j {
		mid := j - (j-i)/2
		if nums[mid] < target {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	left := j

	return right - left - 1
}
