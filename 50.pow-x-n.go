package leetcode

/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	if n == 1 {
		return x
	}
	var a float64 = 1
	var b float64 = x
	for i := n; i > 0; {
		if i%2 == 1 {
			a = b * a
			i--
			continue
		}
		b = b * b
		if i == 2 {
			break
		}
		i /= 2
	}

	return a * b
}

// @lc code=end
