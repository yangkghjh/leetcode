package leetcode

/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	hash := make(map[int]bool)
	ans := []int{}

	for _, v := range nums1 {
		hash[v] = true
	}

	for _, v := range nums2 {
		if hash[v] == true {
			ans = append(ans, v)
			hash[v] = false
		}
	}

	return ans
}

// 时间复杂度： O(N)
// 空间复杂度： O(len(nums1))

// @lc code=end
