package offer

// 在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。

func firstUniqChar(s string) byte {
	hash := make([]int, 26)

	for _, b := range []byte(s) {
		hash[b-'a']++
	}

	for _, b := range []byte(s) {
		if hash[b-'a'] == 1 {
			return b
		}
	}

	return ' '
}
