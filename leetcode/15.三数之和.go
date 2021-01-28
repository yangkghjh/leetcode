package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	length := len(nums)
	results := [][]int{}
	sorted := sort.IntSlice(nums)
	sorted.Sort()
	nums = sorted
	var target int
	var y int
	var z int
	for x := 0; x < length-2; x++ {
		if nums[x] > 0 {
			break
		}
		if x > 0 && nums[x] == nums[x-1] {
			continue
		}
		target = 0 - nums[x]
		y = x + 1
		z = length - 1
		for y < z {
			if nums[y]+nums[z] > target {
				z--
			} else if nums[y]+nums[z] < target {
				y++
			} else {
				results = append(results, []int{nums[x], nums[y], nums[z]})
				y++
				for ; y < z && nums[y] == nums[y-1]; y++ {
				}
			}
		}
	}
	return results
}

// @lc code=end
