package leetcode

/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */

// @lc code=start
// 递归写法
func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	if numRows == 1 {
		return [][]int{{1}}
	}

	row := make([]int, numRows)
	result := generate(numRows - 1)
	prev := result[len(result)-1]

	row[0] = 1
	row[numRows-1] = 1

	for i := 1; i < numRows-1; i++ {
		row[i] = prev[i-1] + prev[i]
	}

	result = append(result, row)

	return result
}

func generateCommon(numRows int) [][]int {
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
