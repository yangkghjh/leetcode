package leetcode

/*
 * @lc app=leetcode.cn id=220 lang=golang
 *
 * [220] 存在重复元素 III
 */

// @lc code=start
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if t < 0 {
		return false
	}
	buckets := make(map[int]int)
	w := t + 1

	genKey := func(x int) int {
		if x < 0 {
			return (x+1)/w - 1
		} else {
			return x / w
		}
	}

	abs := func(data int) int {
		if data < 0 {
			return -1 * data
		}
		return data
	}

	for i := range nums {
		key := genKey(nums[i])
		if _, ok := buckets[key]; ok {
			return true
		}

		if val, ok := buckets[key-1]; ok && abs(nums[i]-val) < w {
			return true
		}
		if val, ok := buckets[key+1]; ok && abs(nums[i]-val) < w {
			return true
		}

		buckets[key] = nums[i]
		if i >= k {
			delete(buckets, genKey(nums[i-k]))
		}
	}

	return false
}

// 时间复杂度： O(N)
// 空间复杂度： O(min(n,k))

// @lc code=end
