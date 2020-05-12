package leetcode

/*
 * @lc app=leetcode.cn id=461 lang=golang
 *
 * [461] 汉明距离
 */

// @lc code=start
func hammingDistance(x int, y int) int {
	z := x ^ y
	res := 0
	for i := 0; i < 32; i++ {
		if z >> i & 1 == 1 {
			res++
		}
	}

	return res
}
// @lc code=end

