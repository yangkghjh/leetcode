package leetcode

/*
 * @lc app=leetcode.cn id=249 lang=golang
 *
 * [249] 移位字符串分组
 */

// @lc code=start
func groupStrings(strings []string) [][]string {
	hash := map[string][]string{}

	for _, s := range strings {
		h := groupStringsHash(s)
		strs, ok := hash[h]

		if !ok {
			strs = []string{s}
			hash[h] = strs
		} else {
			hash[h] = append(strs, s)
		}
	}

	ans := [][]string{}
	for _, strs := range hash {
		ans = append(ans, strs)
	}

	return ans
}

func groupStringsHash(s string) string {
	ans := make([]byte, len(s))

	ss := []byte(s)
	first := ss[0]
	diff := ss[0] - 'a'
	for i, c := range ss {
		if c < first {
			ans[i] = c + (26 - diff)
		} else {
			ans[i] = c - diff
		}
	}

	return string(ans)
}

// @lc code=end
