package leetcode

/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] 反转字符串
 */

// @lc code=start
// 递归写法
func reverseString(s []byte) {
	reverseStringHelper(s, 0, len(s)-1)
}

func reverseStringHelper(s []byte, start, end int) {
	if start < end {
		s[start], s[end] = s[end], s[start]

		reverseStringHelper(s, start+1, end-1)
	}
}

// 时间复杂度： O(N)
// 空间复杂度： O(N)

// func reverseString(s []byte) {
// 	start := 0
// 	end := len(s) - 1

// 	for start < end {
// 		s[start], s[end] = s[end], s[start]
// 		start++
// 		end--
// 	}
// }

// @lc code=end
