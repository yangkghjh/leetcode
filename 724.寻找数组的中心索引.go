package leetcode

/*
 * @lc app=leetcode.cn id=724 lang=golang
 *
 * [724] 寻找数组的中心索引
 */

// @lc code=start
func pivotIndex(nums []int) int {
	// len(nums) [0, 10000]
	// nums[i] [-1000, 1000]
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return 0
	}

	sumleft, sumright := 0, 0
	for _, n := range nums {
		sumright += n
	}
	sumright -= nums[0]

	if sumleft == sumright {
		return 0
	}

	for i := 1; i < len(nums); i++ {
		sumright -= nums[i]
		sumleft += nums[i-1]
		if sumleft == sumright {
			return i
		}
	}

	return -1
}

// @lc code=end
