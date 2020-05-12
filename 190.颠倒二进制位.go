package leetcode

/*
 * @lc app=leetcode.cn id=190 lang=golang
 *
 * [190] 颠倒二进制位
 */

// @lc code=start
func reverseBits(num uint32) uint32 {
	var res uint32
	
	for i := 0; i < 32; i++ {
		res += (num >> i) & 1 << (31 - i)
	}

	return res
}
// @lc code=end
