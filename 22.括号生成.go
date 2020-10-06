package leetcode

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	results := []string{}
	return backtrack(results, "", 0, 0, n)
}

func backtrack(results []string, stack string, open int, close int, max int) []string {
	if len(stack) == max*2 {
		results = append(results, stack)
		return results
	}

	if open < max {
		results = backtrack(results, stack+"(", open+1, close, max)
	}
	if close < open {
		results = backtrack(results, stack+")", open, close+1, max)
	}
	return results
}

// @lc code=end
