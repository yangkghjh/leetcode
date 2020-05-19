package leetcode

/*
 * @lc app=leetcode.cn id=494 lang=golang
 *
 * [494] 目标和
 */

// @lc code=start
func findTargetSumWays(nums []int, S int) int {
	return findTargetSumWaysDFS(nums, 0, S, true) + findTargetSumWaysDFS(nums, 0, S, false)
}

func findTargetSumWaysDFS(nums []int, i, s int, operate bool) int {
	if operate {
		s -= nums[i]
	} else {
		s += nums[i]
	}
	i++

	if i == len(nums) {
		if s == 0 {
			return 1
		} else {
			return 0
		}
	}

	return findTargetSumWaysDFS(nums, i, s, true) + findTargetSumWaysDFS(nums, i, s, false)
}

// @lc code=end
