package leetcode

/*
 * @lc app=leetcode.cn id=91 lang=golang
 *
 * [91] 解码方法
 */

// @lc code=start
func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	ret := 0
	if !numDecodingsChain(s, &ret) {
		return 0
	}

	return ret
}

func numDecodingsChain(s string, ret *int) bool {
	if len(s) == 0 {
		*ret++
		return true
	}
	if s[0] == '0' {
		return true
	}

	if len(s) >= 2 && s[1] == '0' {
		if s[0] > '2' {
			return false
		}
		return numDecodingsChain(s[2:], ret)
	}

	if len(s) >= 2 && s[0] == '2' && s[1] <= '6' {
		if !numDecodingsChain(s[2:], ret) {
			return false
		}
	}
	if len(s) >= 2 && s[0] == '1' {
		if !numDecodingsChain(s[2:], ret) {
			return false
		}
	}

	return numDecodingsChain(s[1:], ret)
}

// @lc code=end
