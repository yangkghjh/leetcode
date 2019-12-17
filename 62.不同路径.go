package leetcode

/*
 * @lc app=leetcode.cn id=62 lang=golang
 *
 * [62] 不同路径
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	var p1 int64 = 1
	var p2 int64 = 1
	long := n - 1
	short := m - 1
	if m > n {
		long = m - 1
		short = n - 1
	}
	for short > 0 {
		p1 *= int64(short)
		p2 *= int64(long) + int64(short)
		short--
	}
	return int(p2 / p1)
}

// @lc code=end
