package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	if right < 0 {
		return math.MinInt64
	}

	for left < right {
		mid := (left + right) / 2

		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return nums[left]
}

// 时间复杂度： O(logn)
// 空间复杂度： O(1)

// @lc code=end
