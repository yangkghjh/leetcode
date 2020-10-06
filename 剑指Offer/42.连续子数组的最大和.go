package offer

// 输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。
// 要求时间复杂度为O(n)。

func maxSubArray(nums []int) int {
	max, sum := nums[0], 0

	for _, num := range nums {
		if sum < 0 {
			sum = num
		} else {
			sum += num
		}

		if sum > max {
			max = sum
		}
	}

	return max
}
