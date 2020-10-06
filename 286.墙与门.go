package leetcode

/*
 * @lc app=leetcode.cn id=284 lang=golang
 *
 * [286] 墙与门
 */

// @lc code=start
func wallsAndGates(rooms [][]int) {
	const INF = 2147483647

	height := len(rooms)
	if height == 0 {
		return
	}
	width := len(rooms[0])

	queue := [][]int{}
	// 找到所有的 0，加入队列
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if rooms[i][j] == 0 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for n := 0; n < len(queue); n++ {
		i, j := queue[n][0], queue[n][1]
		v := rooms[i][j] + 1

		for _, dir := range dirs {
			ni, nj := i+dir[0], j+dir[1]

			if ni >= 0 && ni < height && nj >= 0 && nj < width && rooms[ni][nj] != -1 {
				if rooms[ni][nj] > v {
					rooms[ni][nj] = v
					queue = append(queue, []int{ni, nj})
				}
			}
		}
	}
}

// @lc code=end
