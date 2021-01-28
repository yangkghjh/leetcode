package leetcode

/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	left, right := 0, num

	for left <= right {
		mid := (left + right) / 2
		x := mid * mid
		if x == num {
			return true
		}

		if num < x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}

// @lc code=end
