package leetcode

/*
 * @lc app=leetcode.cn id=119 lang=golang
 *
 * [119] 杨辉三角 II
 */

// @lc code=start
func getRow(rowIndex int) []int {
	ans := make([]int, rowIndex+1)
	ans[0] = 1
	for i := 1; i <= rowIndex; i++ {
		ans[i] = 1
		for j := i; j > 1; j-- {
			ans[j-1] = ans[j-2] + ans[j-1]
		}
	}

	return ans
}

// 时间复杂度 O(n2)
// 空间复杂度 O(k)

// @lc code=end
