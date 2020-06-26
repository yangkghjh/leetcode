package leetcode

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	max := 0
	cur := 0
	hash := map[rune]bool{}
	left := 0

	ss := []rune(s)

	for _, c := range s {
		if _, ok := hash[c]; ok {
			if cur > max {
				max = cur
			}
			for {
				cur--
				if ss[left] == c {
					break
				}
				delete(hash, ss[left])
				left++
			}
			left++
		}

		cur++
		hash[c] = true
	}

	if cur > max {
		return cur
	}

	return max
}

// 时间复杂度： O(N)
// 空间复杂度： O(N)

// @lc code=end
