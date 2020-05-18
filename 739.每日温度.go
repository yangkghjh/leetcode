package leetcode

/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 */

// @lc code=start
func dailyTemperatures(T []int) []int {
	stack := [][]int{}
	res := make([]int, len(T))

	for i, t := range T {
		if len(stack) == 0 {
			stack = append(stack, []int{i, t})
			continue
		}

		for len(stack) >= 1 {
			top := stack[len(stack)-1]
			if top[1] >= t {
				break
			}
			res[top[0]] = i - top[0]
			stack = stack[0 : len(stack)-1]
		}

		stack = append(stack, []int{i, t})
	}

	for _, x := range stack {
		res[x[0]] = 0
	}

	return res
}

// 题解 https://leetcode-cn.com/problems/daily-temperatures/solution/leetcode-tu-jie-739mei-ri-wen-du-by-misterbooo/

// @lc code=end
