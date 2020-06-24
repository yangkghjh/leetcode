package leetcode

/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start
func isHappy(n int) bool {
	hash := map[int]bool{}

	for {
		r := 0
		for n != 0 {
			i := n % 10
			n /= 10
			r += i * i
		}

		if r == 1 {
			return true
		}

		if _, ok := hash[r]; ok {
			return false
		} else {
			hash[r] = true
		}

		n = r
	}
}

// 时间复杂度： O(N)， N 为成为 1 的步骤数
// 空间复杂度： O(N)， N 为成为 1 的步骤数

// @lc code=end
