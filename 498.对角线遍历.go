package leetcode

/*
 * @lc app=leetcode.cn id=498 lang=golang
 *
 * [498] 对角线遍历
 */

// @lc code=start
func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	if len(matrix) == 1 {
		return matrix[0]
	}

	m, n := len(matrix), len(matrix[0])
	res := make([]int, m*n)
	direction := 1

	for i, j := 0, 0; i < m+n-1; i++ {
		x0, y0 := getNextStartPoint(m, n, i)

		if direction == 1 {
			for {
				res[j] = matrix[x0][y0]
				j++
				if x0 == 0 || y0 == n-1 {
					break
				}
				x0--
				y0++
			}
			direction = -1
		} else {
			for {
				res[j] = matrix[x0][y0]
				j++
				if x0 == m-1 || y0 == 0 {
					break
				}
				x0++
				y0--
			}
			direction = 1
		}
	}

	return res
}

func getNextStartPoint(m, n, i int) (int, int) {
	if n == 1 {
		return i, 0
	}
	if i%2 == 0 {
		// 下边沿
		x := i
		if x < m {
			return x, 0
		}
		return m - 1, i - m + 1
	} else {
		// 上边沿
		y := i
		if y < n {
			return 0, y
		}
		return i - n + 1, n - 1
	}
}

// @lc code=end
