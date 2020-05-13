package leetcode

/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */

// @lc code=start
func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for row := 0; row < numRows; row++ {
		r := make([]int, row+1)
		r[0] = 1
		r[row] = 1
		for i := 1; i < row; i++ {
			r[i] = res[row-1][i] + res[row-1][i-1]
		}
		res[row] = r
	}

	return res
}
// @lc code=end

