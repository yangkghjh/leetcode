package leetcode

/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */

// @lc code=start
func setZeroes(matrix [][]int) {
	rows := make([]int, 0)
	cols := make([]int, 0)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				rows = append(rows, i)
				cols = append(cols, j)
			}
		}
	}

	for _, row := range rows {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[row][j] = 0
		}
	}

	for _, col := range cols {
		for i := 0; i < len(matrix); i++ {
			matrix[i][col] = 0
		}
	}
}

// @lc code=end
