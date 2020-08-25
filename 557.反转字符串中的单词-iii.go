package leetcode

/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] 反转字符串中的单词 III
 */

// @lc code=start
func reverseWords(s string) string {
	str := []byte(s)
	wordStart := 0
	l := len(str)
	for i := 0; i <= l; i++ {
		if i == l || str[i] == ' ' {
			for j, k := wordStart, i-1; j < k; j, k = j+1, k-1 {
				str[j], str[k] = str[k], str[j]
			}
			wordStart = i + 1
		}
	}

	return string(str)
}

// 时间复杂度： O(n)
// 空间复杂度： O(1)

// @lc code=end
