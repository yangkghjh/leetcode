package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	length := len(nums)
	results := [][]int{}
	sort.Ints(nums)

	for i := 0; i < length-3; i++ {
		for j := i + 1; j < length-2; j++ {
			value := target - nums[i] - nums[j]

			x := j + 1
			y := length - 1

			for x < y {
				if nums[x]+nums[y] > value {
					y--
				} else if nums[x]+nums[y] < value {
					x++
				} else {
					results = append(results, []int{nums[i], nums[j], nums[x], nums[y]})
					x++
					for ; x < y && nums[x] == nums[x-1]; x++ {
					}
				}
			}

			for j+1 < length-2 && nums[j] == nums[j+1] {
				j++
			}
		}

		for i+1 < length-3 && nums[i] == nums[i+1] {
			i++
		}
	}

	return results
}

// 时间复杂度： O(N3)
// 空间复杂度： O(1)

// @lc code=end
