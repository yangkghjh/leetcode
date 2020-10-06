package leetcode

/*
 * @lc app=leetcode.cn id=485 lang=golang
 *
 * [485] 最大连续1的个数
 */

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			n++
		} else {
			if n > max {
				max = n
			}
			n = 0
		}
	}

	if n > max {
		return n
	}

	return max
}

// 时间复杂度 O(n)
// 空间复杂度 O(1)

// @lc code=end
