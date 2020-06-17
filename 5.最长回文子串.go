package leetcode

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
// 方法二中心扩散法
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	n := len(s)
	start, end := 0, 0
	for i := 0; i < n; i++ {
		start1, end1 := expendAroundCenter(s, i, i)
		start2, end2 := expendAroundCenter(s, i, i+1)

		if end1-start1 > end-start {
			start, end = start1, end1
		}
		if end2-start2 > end-start {
			start, end = start2, end2
		}
	}

	return s[start : end+1]
}

func expendAroundCenter(s string, start, end int) (int, int) {
	for ; start >= 0 && end < len(s) && s[start] == s[end]; start, end = start-1, end+1 {
	}

	return start + 1, end - 1
}

// 方法一 动态规划
func longestPalindromeDynamicProgramming(s string) string {
	ans := ""
	n := len(s)
	// 结构数组 dynamic programming
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}

	for l := 0; l < n; l++ {
		for i := 0; i+l < n; i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = 1
			} else if l == 1 {
				if s[i] == s[j] {
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] == 1 && l+1 > len(ans) {
				ans = s[i : j+1]
			}
		}
	}

	return ans
}

// 时间复杂度 O(n2)
// 空间复杂度 O(n2)

// @lc code=end
