package leetcode

/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	lengthy := len(matrix)
	if lengthy == 0 {
		return []int{}
	}
	lengthx := len(matrix[0])
	total := lengthy * lengthx
	ans := make([]int, total)

	point := []int{0, 0}
	coef := []int{1, -1}
	round := 0
	target := lengthx
	for i := 0; i < total; i++ {
		ans[i] = matrix[point[1]][point[0]]
		target--
		if target == 0 {
			round++

			if round%2 == 0 {
				lengthx--
				target = lengthx
			} else {
				lengthy--
				target = lengthy
			}
		}
		point[round%2] += coef[(round%4)/2]
	}

	return ans
}

// @lc code=end
