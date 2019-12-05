package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := math.MinInt64
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if result < sum {
			result = sum
		}
		if sum < 0 {
			sum = 0
		}
	}

	return result
}

// @lc code=end
