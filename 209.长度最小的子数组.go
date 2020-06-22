package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(s int, nums []int) int {
	ans := math.MaxInt64
	window := 0
	for left, right := 0, 0; right < len(nums); right++ {
		window += nums[right]
		for window >= s {
			if right+1-left < ans {
				ans = right + 1 - left
				if ans == 1 {
					return 1
				}
			}
			window -= nums[left]
			left++
		}
	}

	if ans == math.MaxInt64 {
		return 0
	}
	return ans
}

// 时间复杂度 O(n)
// 空间复杂度 O(1)

// @lc code=end
