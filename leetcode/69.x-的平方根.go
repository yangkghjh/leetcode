package leetcode

/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] x 的平方根
 */

// @lc code=start
func mySqrt(x int) int {
	// func mySqrt(x int) int {
	// 二分法
	left, right := 0, x
	ans := -1
	for right >= left {
		mid := left + (right-left)/2
		if mid*mid <= x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}

func mySqrtNewtonMethod(x int) int {
	// func mySqrt(x int) int {
	// 牛顿法
	if x == 0 {
		return 0
	}

	var last float64 = 0
	var res float64 = 1
	for res != last {
		last = res
		res = (res + float64(x)/res) / 2
	}
	return int(res)
}

// @lc code=end
