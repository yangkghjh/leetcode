package leetcode

/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] 旋转图像
 */

// @lc code=start
func rotate(matrix [][]int) {
	if len(matrix) <= 1 {
		return
	}
	// 确认旋转区域
	l := len(matrix)
	ymax := l/2 + l%2
	xmax := l / 2

	// 旋转可以分解为左移和下移
	stack := 0
	for x := 0; x < xmax; x++ {
		for y := 0; y < ymax; y++ {
			stack = matrix[y][x]
			matrix[y][x] = matrix[l-x-1][y]
			matrix[l-x-1][y] = matrix[l-y-1][l-x-1]
			matrix[l-y-1][l-x-1] = matrix[x][l-y-1]
			matrix[x][l-y-1] = stack
		}
	}
}

// @lc code=end
