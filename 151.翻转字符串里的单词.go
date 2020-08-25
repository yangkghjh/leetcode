package leetcode

import (
	"bytes"
	"strings"
)

/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 翻转字符串里的单词
 */

// @lc code=start
func reverseWords1(s string) string {
	strs := strings.Split(s, " ")

	ans := bytes.NewBuffer([]byte{})
	for i := len(strs) - 1; i >= 0; i-- {
		if strs[i] != "" {
			ans.WriteString(" ")
			ans.WriteString(strs[i])
		}
	}

	if ans.Len() == 0 {
		return ""
	}

	return ans.String()[1:]
}

// 时间复杂度 O(n)
// 空间复杂度 O(n)

// @lc code=end
