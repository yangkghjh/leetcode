package leetcode

/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 报数
 */

// @lc code=start
func countAndSay(n int) string {
	return string(say(n, []byte{'1'}))
}

func say(n int, s []byte) []byte {
	if n == 1 {
		return s
	}

	var result []byte
	stack := s[0]
	var num byte
	for _, c := range s {
		if c == stack {
			num++
		} else {
			result = append(result, num+48)
			result = append(result, stack)
			stack = c
			num = 1
		}
	}
	if num != 0 {
		result = append(result, num+48)
		result = append(result, stack)
	}

	n--
	return say(n, result)
}

// @lc code=end
