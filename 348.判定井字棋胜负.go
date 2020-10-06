package leetcode

/*
 * @lc app=leetcode.cn id=348 lang=golang
 *
 * [348] 判定井字棋胜负
 */

// @lc code=start
type TicTacToe struct {
	stack [][]int
	size  int
}

/** Initialize your data structure here. */
func TicTacToeConstructor(n int) TicTacToe {
	return TicTacToe{
		stack: [][]int{nil, make([]int, 2*n+2), make([]int, 2*n+2)},
		size:  n,
	}
}

/** Player {player} makes a move at ({row}, {col}).
  @param row The row of the board.
  @param col The column of the board.
  @param player The player, can be either 1 or 2.
  @return The current winning condition, can be either:
          0: No one wins.
          1: Player 1 wins.
          2: Player 2 wins. */
func (this *TicTacToe) Move(row int, col int, player int) int {
	items := getTicTacToeStackID(row, col, this.size)

	for _, id := range items {
		this.stack[player][id]++
		if this.stack[player][id] == this.size {
			return player
		}
	}

	return 0
}

func getTicTacToeStackID(row, col, size int) []int {
	if row == col {
		if row+col == size-1 {

			return []int{row, size + col, size * 2, size*2 + 1}
		}
		return []int{row, size + col, size * 2}
	} else if row+col == size-1 {
		return []int{row, size + col, size*2 + 1}
	}

	return []int{row, size + col}
}

/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */
// @lc code=end
