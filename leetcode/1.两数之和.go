package leetcode

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	hash := map[int]int{}

	for key, num := range nums {
		v := target - num
		if k, ok := hash[v]; ok {
			return []int{k, key}
		}
		hash[num] = key
	}

	return []int{}
}

// 时间复杂度： O(n)
// 空间复杂度： O(n)

// @lc code=end
