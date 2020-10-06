package leetcode

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	stack := []byte{}

	for _, mark := range []byte(s) {
		if isStartMark(mark) {
			// push
			stack = append(stack, mark)
		} else {
			if len(stack) == 0 {
				return false
			}
			if isMatch(stack[len(stack)-1], mark) {
				// pop
				stack = stack[0 : len(stack)-1]
			} else {
				return false
			}
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}

// 是否为起始符
func isStartMark(mark byte) bool {
	if mark == '(' || mark == '{' || mark == '[' {
		return true
	}

	return false
}

// 是否匹配
func isMatch(startMark, endMark byte) bool {
	if startMark == '{' && endMark == '}' {
		return true
	}
	if startMark == '(' && endMark == ')' {
		return true
	}
	if startMark == '[' && endMark == ']' {
		return true
	}
	return false
}

// @lc code=end
