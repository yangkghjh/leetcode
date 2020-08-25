package offer

// 输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。

func spiralOrder(matrix [][]int) []int {
	maxY, minY := len(matrix), 0
	if maxY == 0 {
		return []int{}
	}
	maxX, minX := len(matrix[0]), 0

	const right, down, left, up = 0, 1, 2, 3
	direction := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	x, y := 0, 0
	ans := []int{}
	d := right

	for minX < maxX && minY < maxY {
		ans = append(ans, matrix[y][x])

		x1, y1 := x+direction[d][1], y+direction[d][0]
		if x1 < minX || x1 >= maxX || y1 < minY || y1 >= maxY {
			if d == right {
				minY++
			} else if d == down {
				maxX--
			} else if d == left {
				maxY--
			} else if d == up {
				minX++
			}

			d++
			if d > 3 {
				d = 0
			}

			x1, y1 = x+direction[d][1], y+direction[d][0]
		}

		x, y = x1, y1
	}

	return ans
}
