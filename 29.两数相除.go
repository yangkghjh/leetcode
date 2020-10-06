package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] 两数相除
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	if divisor == 0 || dividend == 0 {
		return 0
	}

	divisorT0 := true
	dividendT0 := true
	if divisor < 0 {
		divisorT0 = false
		divisor = -divisor
	}
	if dividend < 0 {
		dividendT0 = false
		dividend = -dividend
	}
	var result int
	if (divisorT0 && !dividendT0) || (!divisorT0 && dividendT0) {
		result = -leftMove(dividend, divisor)
	} else {
		result = leftMove(dividend, divisor)
	}

	if result > int(math.MaxInt32) || result < int(math.MinInt32) {
		return int(math.MaxInt32)
	}

	return result
}

func leftMove(dividend, divisor int) int {
	result := 1
	com := divisor
	if dividend < divisor {
		return 0
	}
	for {
		com = com << 1
		if com > dividend {
			com = com >> 1
			if dividend-com >= divisor {
				return leftMove(dividend-com, divisor) + result
			}
			return result
		}
		result = result << 1
	}
}

// @lc code=end
