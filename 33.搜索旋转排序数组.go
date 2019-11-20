package leetcode

// import "fmt"

/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if nums[0] == target {
		return 0
	}
	if nums[len(nums)-1] == target {
		return len(nums) - 1
	}

	return binarySearch(nums, target, 0, len(nums)-1)
}

func binarySearch(nums []int, target, start, end int) int {
	mid := (start + end) / 2
	if nums[mid] == target {
		return mid
	}

	// fmt.Println(start, end)
	if end-start <= 1 {
		return -1
	}

	if nums[start] < nums[mid] { // 前半段有序
		if target < nums[mid] && target > nums[start] { // 在前半段
			return binarySearch(nums, target, start, mid)
		}
		// 不在前半段
		return binarySearch(nums, target, mid, end)
	} else {
		if target > nums[mid] && target < nums[end] { // 在后半段
			return binarySearch(nums, target, mid, end)
		}
		// 不在后半段
		return binarySearch(nums, target, start, mid)
	}
}

// @lc code=end
