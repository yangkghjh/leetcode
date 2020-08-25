package leetcode

/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
func sortColors(nums []int) {
	var red int
	var blue int
	l := len(nums)

	for i := 0; i < len(nums); {
		if nums[i] == 0 {
			red++
			nums[i] = nums[red-1]
			nums[red-1] = 0
			if i < red {
				i++
			}
		} else if nums[i] == 2 {
			if i >= l-blue {
				break
			}
			blue++
			nums[i] = nums[l-blue]
			nums[l-blue] = 2
		} else {
			i++
		}
	}
}

// @lc code=end
