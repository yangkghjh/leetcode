package leetcode

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
func numSquares(n int) int {
	queue := []int{n, -1}

	res := 1
	for o := 0; o < len(queue); o++ {
		t := queue[o]
		if t == -1 {
			if o+1 == len(queue) {
				break
			}
			res++
			queue = append(queue, -1)
		}
		max := int(math.Sqrt(float64(t)))
		for i := max; i > 0; i-- {
			x := t - i*i
			if x == 0 {
				return res
			}
			queue = append(queue, x)
		}
	}

	return -1
}

// @lc code=end
