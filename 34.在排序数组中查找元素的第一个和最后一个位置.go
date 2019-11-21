package leetcode

/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	return binarySearchRange(nums, target, 0, len(nums)-1)
}

func binarySearchRange(nums []int, target, start, end int) []int {
	mid := (start + end) / 2
	if nums[mid] == target {
		result := []int{mid, mid}
		for min := mid; min >= 0 && nums[min] == target; min-- {
			result[0] = min
		}
		for max := mid; max < len(nums) && nums[max] == target; max++ {
			result[1] = max
		}
		return result
	}
	if end-start <= 1 {
		if nums[end] == target {
			return []int{end, end}
		}
		return []int{-1, -1}
	}

	if nums[mid] > target {
		return binarySearchRange(nums, target, start, mid)
	}
	return binarySearchRange(nums, target, mid, end)
}

// @lc code=end
