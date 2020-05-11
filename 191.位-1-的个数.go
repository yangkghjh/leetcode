package leetcode

/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */

// @lc code=start
func hammingWeight(num uint32) int {
	res := 0
	var i uint32
    for ; i < 32; i++ {
		if (num >> i & 1) == 1 {
			res++
		}
	}

	return res
}
// @lc code=end

