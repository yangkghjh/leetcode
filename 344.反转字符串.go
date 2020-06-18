package leetcode

/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] 反转字符串
 */

// @lc code=start
func reverseString(s []byte) {
	start := 0
	end := len(s) - 1

	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}

// @lc code=end
