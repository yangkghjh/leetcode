package leetcode

import (
	"math"
	"sort"
)

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	length := len(nums)
	var x, y, z int

	var sum int
	s := sort.IntSlice(nums)
	s.Sort()
	nums = s
	result := nums[0] + nums[1] + nums[2]

	for ; x < length-2; x++ {
		y = x + 1
		z = length - 1

		for y < z {
			sum = nums[x] + nums[y] + nums[z]
			if math.Abs(float64(result-target)) > math.Abs(float64(sum-target)) {
				result = sum
			}
			// fmt.Println(x, y, z, sum, math.Abs(float64(result-target)), math.Abs(float64(sum-target)), result)
			if sum > target {
				z--
			} else if sum < target {
				y++
			} else {
				return result
			}
		}
	}

	return result
}

// @lc code=end
