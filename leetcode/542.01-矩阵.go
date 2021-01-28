package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=542 lang=golang
 *
 * [542] 01 矩阵
 */

// @lc code=start
func updateMatrix(matrix [][]int) [][]int {

	height := len(matrix)
	if height == 0 {
		return [][]int{}
	}
	width := len(matrix[0])
	if width == 0 {
		return [][]int{}
	}

	queue := [][]int{}
	ans := make([][]int, height)
	for i := 0; i < height; i++ {
		ans[i] = make([]int, width)
		for j := 0; j < width; j++ {
			if matrix[i][j] == 0 {
				queue = append(queue, []int{i, j})
				ans[i][j] = 0
			} else {
				ans[i][j] = math.MaxInt32
			}
		}
	}

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for x := 0; x < len(queue); x++ {
		i, j := queue[x][0], queue[x][1]
		v := ans[i][j]

		for _, d := range dirs {
			ni, nj := i+d[0], j+d[1]
			if ni >= 0 && nj >= 0 && ni < height && nj < width && ans[ni][nj] == math.MaxInt32 {
				ans[ni][nj] = v + 1
				queue = append(queue, []int{ni, nj})
			}
		}
	}

	return ans
}

// @lc code=end
