package leetcode

/*
 * @lc app=leetcode.cn id=287 lang=golang
 *
 * [287] 寻找重复数
 */

// @lc code=start

func findDuplicate(nums []int) int {
	n := len(nums)
	l, r := 1, n-1
	ans := -1
	for l < r {
		mid := (l + r) / 2

		cnt := 0
		for i := 0; i < n; i++ {
			if nums[i] <= mid {
				cnt++
			}
		}

		if cnt > mid {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return ans
}

// @lc code=end
