package leetcode

import (
	"sort"
	"strings"
)

/*
 * @lc app=leetcode.cn id=648 lang=golang
 *
 * [648] 单词替换
 */

// @lc code=start
func replaceWords(dict []string, sentence string) string {
	sort.Slice(dict, func(i, j int) bool {
		return len(dict[i]) < len(dict[j])
	})

	words := strings.Split(sentence, " ")
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(dict); j++ {
			if strings.HasPrefix(words[i], dict[j]) {
				words[i] = dict[j]
				break
			}
		}
	}

	return strings.Join(words, " ")
}

// func replaceWords(dict []string, sentence string) string {
// 	trie := &ReplaceWordsTrie{
// 		Children: map[byte]*ReplaceWordsTrie{},
// 	}

// 	for _, prefix := range dict {
// 		trie.Insert(prefix)
// 	}

// 	words := strings.Split(sentence, " ")
// 	ans := ""

// 	for _, word := range words {
// 		prefix := trie.Search(word)

// 		if prefix == "" {
// 			prefix = word
// 		}

// 		ans = ans + prefix + " "
// 	}

// 	if len(ans) == 0 {
// 		return ans
// 	}

// 	return ans[:len(ans)-1]
// }

type ReplaceWordsTrie struct {
	Children map[byte]*ReplaceWordsTrie
	Finish   bool
}

func (r *ReplaceWordsTrie) Insert(prefix string) {
	node := r

	for _, l := range []byte(prefix) {
		n, ok := node.Children[l]

		if !ok {
			n = &ReplaceWordsTrie{
				Children: map[byte]*ReplaceWordsTrie{},
			}
			node.Children[l] = n
		}

		node = n
	}

	node.Finish = true
}

func (r *ReplaceWordsTrie) Search(word string) string {
	node := r
	w := []byte(word)
	for i := range w {
		n, ok := node.Children[w[i]]

		if !ok {
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
