package leetcode

/*
 * @lc app=leetcode.cn id=658 lang=golang
 *
 * [658] 找到 K 个最接近的元素
 */

// @lc code=start
func findClosestElements(arr []int, k int, x int) []int {
	l := len(arr)
	left, right := 0, l-k

	for left < right {
		mid := (left + right) / 2

		if x-arr[mid] > arr[mid+k]-x {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return arr[left : left+k]
}

// @lc code=end
