package leetcode

/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 缺失数字
 */

// @lc code=start
func missingNumber(nums []int) int {
	n := len(nums) + 1

	xor := 0
	for i := 1; i < n; i++ {
		xor ^= i
	}

	for i := 0; i < n-1; i++ {
		xor ^= nums[i]
	}

	return xor
}
// @lc code=end

