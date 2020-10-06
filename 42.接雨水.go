package leetcode

/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	// 双指针法
	if len(height) <= 1 {
		return 0
	}
	left := 0
	right := len(height) - 1
	result := 0
	max := 0  // 高度
	mode := 0 // 0 左缩进，1 右缩进

	for {
		if left >= right {
			return result
		}

		if mode == 0 {
			if height[left] <= max {
				result += max - height[left]
				left++
			} else {
				mode = 1
			}
		} else {
			if height[right] <= max {
				result += max - height[right]
				right--
			} else {
				mode = 0
				max++
			}
		}
	}
}

// @lc code=end
