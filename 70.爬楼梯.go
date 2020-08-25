package leetcode

/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	memo := make(map[int]int, n)
	return climbStairsSteps(0, n, memo)
}

func climbStairsSteps(i int, n int, memo map[int]int) int {
	if i > n {
		return 0
	}

	if i == n {
		return 1
	}

	if n, ok := memo[i]; ok {
		return n
	}

	memo[i] = climbStairsSteps(i+1, n, memo) + climbStairsSteps(i+2, n, memo)
	return memo[i]
}

// @lc code=end
