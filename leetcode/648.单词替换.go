package leetcode

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=648 lang=golang
 *
 * [648] 单词替换
 */

// @lc code=start
// func replaceWords(dict []string, sentence string) string {
// 	sort.Slice(dict, func(i, j int) bool {
// 		return len(dict[i]) < len(dict[j])
// 	})

// 	words := strings.Split(sentence, " ")
// 	for i := 0; i < len(words); i++ {
// 		for j := 0; j < len(dict); j++ {
// 			if strings.HasPrefix(words[i], dict[j]) {
// 				words[i] = dict[j]
// 				break
// 			}
// 		}
// 	}

// 	return strings.Join(words, " ")
// }

func replaceWords(dict []string, sentence string) string {
	trie := new(ReplaceWordsTrie)

	for _, prefix := range dict {
		trie.Insert(prefix)
	}

	words := strings.Split(sentence, " ")
	ans := ""

	for _, word := range words {
		prefix := trie.Search(word)

		if prefix == "" {
			prefix = word
		}

		ans = ans + prefix + " "
	}

	if len(ans) == 0 {
		return ans
	}

	return ans[:len(ans)-1]
}

type ReplaceWordsTrie struct {
	Finish   bool
	Children [26]*ReplaceWordsTrie
}

func (r *ReplaceWordsTrie) Insert(prefix string) {
	node := r

	for _, l := range []byte(prefix) {
		x := int(l - 'a')
		n := node.Children[x]

		if n == nil {
			n = new(ReplaceWordsTrie)
			node.Children[x] = n
		}

		node = n
	}

	node.Finish = true
}

func (r *ReplaceWordsTrie) Search(word string) string {
	node := r
	w := []byte(word)
	for i := range w {
		x := int(w[i] - 'a')
		n := node.Children[x]

		if n == nil {
			return ""
		}

		if n.Finish {
			return string(w[:i+1])
		}

		node = n
	}

	return ""
}

// @lc code=end
