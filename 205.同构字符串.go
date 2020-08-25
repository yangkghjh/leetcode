package leetcode

import "bytes"

/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	return isomorphicHash(s) == isomorphicHash(t)
}

func isomorphicHash(s string) string {
	w := '1'

	buf := new(bytes.Buffer)
	hash := map[rune]rune{}

	for _, c := range s {
		k, ok := hash[c]
		if !ok {
			hash[c] = w
			k = w
			w++
		}

		buf.WriteRune(k)
	}

	return buf.String()
}

// 时间复杂度： O(N)
// 空间复杂度： O(N)

// @lc code=end
