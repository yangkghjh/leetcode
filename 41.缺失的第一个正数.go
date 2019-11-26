package leetcode

/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	// 桶排序
	for i := 0; i < len(nums); {
		if nums[i]-1 == i || nums[i] <= 0 || nums[i] > len(nums) { // 位置正确或者数值为非正数
			i++
			continue
		}

		// 交换到正确位置
		x := nums[nums[i]-1]
		if x == nums[i] { // 位置上的数据正确
			i++
			continue
		}

		nums[nums[i]-1] = nums[i]
		nums[i] = x
	}

	// 检查异常值
	for i := 0; i < len(nums); i++ {
		if nums[i]-1 != i {
			return i + 1
		}
	}
	return nums[len(nums)-1] + 1
}

// @lc code=end
