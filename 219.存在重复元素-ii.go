package leetcode

/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	hash := map[int]int{}

	for p2, v := range nums {
		p1, ok := hash[v]
		if ok && p2-p1 <= k {
			return true
		}

		hash[v] = p2
	}

	return false
}

// 时间复杂度： O(N)
// 空间复杂度： O(N)

// @lc code=end
