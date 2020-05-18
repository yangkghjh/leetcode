package leetcode

import "strconv"

/*
 * @lc app=leetcode.cn id=150 lang=golang
 *
 * [150] 逆波兰表达式求值
 */

// @lc code=start
func evalRPN(tokens []string) int {
	stack := []int{}

	for _, t := range tokens {
		n, err := strconv.Atoi(t)
		if err != nil {
			var r int
			x := stack[len(stack)-2]
			y := stack[len(stack)-1]
			if t == "+" {
				r = x + y
			} else if t == "-" {
				r = x - y
			} else if t == "*" {
				r = x * y
			} else if t == "/" {
				r = x / y
			}
			stack[len(stack)-2] = r
			stack = stack[0 : len(stack)-1]
		} else {
			stack = append(stack, n)
		}
	}

	return stack[0]
}

// @lc code=end
