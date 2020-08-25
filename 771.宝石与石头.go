package leetcode

/*
 * @lc app=leetcode.cn id=771 lang=golang
 *
 * [771] 宝石与石头
 */

// @lc code=start
func numJewelsInStones(J string, S string) int {
	hash := map[rune]bool{}

	for _, j := range J {
		hash[j] = true
	}

	ans := 0

	for _, s := range S {
		if _, ok := hash[s]; ok {
			ans++
		}
	}

	return ans
}

// 时间复杂度： O(len(J)+len(S))
// 空间复杂度： O(len(J))

// @lc code=end
