package leetcode

/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1
	x := m + n - 1
	for i >= -1 && j >= 0 {
		if i == -1 || nums1[i] < nums2[j] {
			nums1[x] = nums2[j]
			j--
		} else {
			nums1[x] = nums1[i]
			i--
		}
		x--
	}
}

// @lc code=end
