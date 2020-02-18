package leetcode

/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	ret := [][]int{[]int{}}
	for i := 0; i < len(nums); i++ {
		for _, sub := range ret {
			sub = append(sub, nums[i])
			tmp := append([]int{}, sub...)
			ret = append(ret, tmp)
		}
	}
	return ret
}

// @lc code=end
