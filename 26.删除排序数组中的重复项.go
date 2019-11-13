package main

/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	x, y := 0, 1
	for ; y < len(nums); y++ {
		if nums[x] != nums[y] {
			x++
			nums[x] = nums[y]
		}

	}
	return x + 1
}

// @lc code=end
