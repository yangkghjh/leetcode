package leetcode

/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lc code=start
func twoSum2(numbers []int, target int) []int {
	index1, index2 := 0, len(numbers)-1

	for {
		if numbers[index1]+numbers[index2] == target {
			return []int{index1 + 1, index2 + 1}
		} else if numbers[index1]+numbers[index2] < target {
			index1++
		} else {
			index2--
		}
	}
}

// @lc code=end
