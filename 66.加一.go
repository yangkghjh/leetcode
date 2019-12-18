package leetcode

/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	item := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = digits[i] + 1
		if digits[i] == 10 {
			digits[i] = 0
		} else {
			item = 0
			break
		}
	}

	if item == 1 {
		digits = append([]int{1}, digits...)
	}

	return digits
}

// @lc code=end
