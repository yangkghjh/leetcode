package leetcode

/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrtBinarySearch(x int) int {
	// func mySqrt(x int) int {
	// 二分法
	m0 := 0
	m1 := x
	m := 0
	for {
		m = (m0 + m1) / 2
		if m*m <= x {
			if m*m == x {
				return m
			}
			if (m+1)*(m+1) >= x {
				if (m+1)*(m+1) == x {
					return m + 1
				}
				return m
			}

			m0 = m
		} else {
			if (m-1)*(m-1) <= x {
				return m - 1
			}

			m1 = m
		}
	}
}

// func mySqrtNewtonMethod(x int) int {
func mySqrt(x int) int {
	// 牛顿法
	if x == 0 {
		return 0
	}

	var last float64 = 0
	var res float64 = 1
	for res != last {
		last = res
		res = (res + float64(x)/res) / 2
	}
	return int(res)
}

// @lc code=end
