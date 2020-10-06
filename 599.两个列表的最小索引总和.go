package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=599 lang=golang
 *
 * [599] 两个列表的最小索引总和
 */

// @lc code=start
func findRestaurant(list1 []string, list2 []string) []string {
	hash := map[string]int{}
	for k, v := range list1 {
		hash[v] = k
	}

	ans := []string{}
	min := math.MaxInt64

	for k, v := range list2 {
		if c, ok := hash[v]; ok {
			if c+k < min {
				ans = []string{v}
				min = c + k
			} else if c+k == min {
				ans = append(ans, v)
			}
		}
	}

	return ans
}

// 时间复杂度： O(N1+N2)
// 空间复杂度： O(N1)

// @lc code=end
