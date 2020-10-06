package leetcode

/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	last := len(nums) - 1
	for i := last - 1; i >= 0; i-- {
		if i+nums[i] >= last {
			last = i
		}
	}
	return last == 0
}

// @lc code=end
