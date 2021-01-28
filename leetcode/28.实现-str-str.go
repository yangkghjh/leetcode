package leetcode

/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	lenght := len(haystack) - len(needle)
	if lenght < 0 {
		return -1
	}
	targetLength := len(needle)

	for i := 0; i <= lenght; i++ {
		if haystack[i:i+targetLength] == needle {
			return i
		}
	}

	return -1
}

// @lc code=end
