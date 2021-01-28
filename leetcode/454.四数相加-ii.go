package leetcode

/*
 * @lc app=leetcode.cn id=454 lang=golang
 *
 * [454] 四数相加 II
 */

// @lc code=start
func fourSumCount(A []int, B []int, C []int, D []int) int {
	hash := map[int]int{}

	for _, i := range A {
		for _, j := range B {
			x := i + j
			_, ok := hash[x]
			if ok {
				hash[x]++
			} else {
				hash[x] = 1
			}
		}
	}

	ans := 0
	for _, i := range C {
		for _, j := range D {
			x := 0 - i - j
			_, ok := hash[x]
			if ok {
				ans += hash[x]
			}
		}
	}
	return ans
}

// 时间复杂度： O(N2)
// 空间复杂度： O(N2

// @lc code=end
