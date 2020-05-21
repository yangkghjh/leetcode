package leetcode

/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */

// @lc code=start
func decodeString(s string) string {
	res, _ := decodeStringInSinglePear(s, 0)
	return res
}

func decodeStringInSinglePear(s string, pos int) (string, int) {
	res := ""
	num := 0
	i := pos

	for ; i < len(s); i++ {
		k := s[i]
		if k >= '0' && k <= '9' {
			num = num*10 + int(k-'0')
		} else if k == '[' {
			tmp, n := decodeStringInSinglePear(s, i+1)
			i = n
			for j := 0; j < num; j++ {
				res += tmp
			}
			num = 0
		} else if k == ']' {
			return res, i
		} else {
			res += string(k)
		}
	}

	return res, i
}

// @lc code=end
