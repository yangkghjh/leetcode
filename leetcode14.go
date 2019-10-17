package main

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	length := len(strs[0])
	for _, str := range strs {
		if len(str) < length {
			length = len(str)
		}
	}

	var prefix string
	for i := 0; i < length; i++ {
		prefix = strs[0][0 : i+1]
		for _, str := range strs[1:] {
			if prefix != str[0:i+1] {
				return prefix[0:i]
			}
		}
	}

	return prefix
}

// @lc code=end
