package leetcode

/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	boards = board
	words = []byte(word)
	lenx, leny = len(board), len(board[0])
	visited = make([][]bool, lenx)
	for i := 0; i < lenx; i++ {
		visited[i] = make([]bool, leny)
	}
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[0]); y++ {
			if board[x][y] == word[0] {
				visited[x][y] = true
				if continueSearch(x, y, 1) {
					return true
				}
				visited[x][y] = false
			}
		}
	}

	return false
}

var (
	directions = [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	lenx, leny int
	visited    [][]bool
	boards     [][]byte
	words      []byte
)

func continueSearch(x, y, index int) bool {
	if index == len(words) {
		return true
	}
	for i := 0; i < 4; i++ {
		x1 := directions[i][0] + x
		y1 := directions[i][1] + y

		if inArray(x1, y1) && boards[x1][y1] == words[index] && !visited[x1][y1] {
			visited[x][y] = true
			if continueSearch(x1, y1, index+1) {
				return true
			}
			visited[x1][y1] = false
		}
	}

	return false
}

func inArray(x, y int) bool {
	return x >= 0 && y >= 0 && x < lenx && y < leny
}

// @lc code=end
