package leetcode

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	results := [][]int{}
	permuteRecursion(&results, nums, 0)
	return results
}

func permuteRecursion(results *[][]int, nums []int, start int) {
	if start == len(nums)-1 {
		*results = append(*results, nums)
		return
	}

	for i := start; i < len(nums); i++ {
		n := make([]int, len(nums))
		copy(n, nums)
		x := n[i]
		n[i] = n[start]
		n[start] = x

		permuteRecursion(results, n, start+1)
	}
}

// @lc code=end
