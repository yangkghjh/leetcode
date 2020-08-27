package offer

import (
	"strings"
)

// 输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。为简单起见，标点符号和普通字母一样处理。例如输入字符串"I am a student. "，则输出"student. a am I"。

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	ans := []string{}

	for i := len(words) - 1; i >= 0; i-- {
		w := strings.TrimSpace(words[i])
		if w != "" {
			ans = append(ans, w)
		}
	}

	return strings.Join(ans, " ")
}
