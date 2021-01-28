package leetcode

/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	lists := make([]map[byte]bool, 9)
	squares := make([]map[byte]bool, 9)
	var i byte
	for ; i < 9; i++ {
		rows[i] = make(map[byte]bool, 9)
		lists[i] = make(map[byte]bool, 9)
		squares[i] = make(map[byte]bool, 9)
	}
	var x byte
	for ; x < 9; x++ {
		var y byte
		for ; y < 9; y++ {
			num := board[y][x]
			if num == '.' {
				continue
			}
			if _, ok := rows[y][num]; ok {
				return false
			}
			rows[y][num] = true

			if _, ok := lists[x][num]; ok {
				return false
			}
			lists[x][num] = true

			s := getSquareSets(x, y)
			if _, ok := squares[s][num]; ok {
				return false
			}
			squares[s][num] = true
		}
	}

	return true
}

// 计算方块组的编号
func getSquareSets(x, y byte) byte {
	return x/3 + y/3*3
}

// @lc code=end
