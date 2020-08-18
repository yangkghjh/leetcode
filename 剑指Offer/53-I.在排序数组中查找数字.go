package offer

// 统计一个数字在排序数组中出现的次数。

func search(nums []int, target int) int {
	ans := 0

	for _, num := range nums {
		if num == target {
			ans++
		} else if ans > 0 {
			break
		}
	}

	return ans
}
