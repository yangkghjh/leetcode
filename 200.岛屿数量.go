package leetcode

/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	res := 0

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '0' {
				continue
			}

			res++
			tagIslandBFS(grid, x, y)
		}
	}

	return res
}

func tagIslandBFS(grid [][]byte, x, y int) {
	if grid[y][x] == '0' {
		return
	}

	grid[y][x] = '0'

	if x > 0 {
		tagIslandBFS(grid, x-1, y)
	}
	if x < len(grid[0])-1 {
		tagIslandBFS(grid, x+1, y)
	}
	if y > 0 {
		tagIslandBFS(grid, x, y-1)
	}
	if y < len(grid)-1 {
		tagIslandBFS(grid, x, y+1)
	}
}

// @lc code=end
