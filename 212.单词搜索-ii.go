package leetcode

/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 */

// @lc code=start
func findWords(board [][]byte, words []string) []string {
	trie := new(FindWordsTrie)
	for _, word := range words {
		trie.Insert(word)
	}

	result := []string{}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			findWordsSearch(i, j, board, trie, &result)
		}
	}

	return result
}

func findWordsSearch(i, j int, board [][]byte, node *FindWordsTrie, result *[]string) {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return
	}

	c := board[i][j]

	if c == '#' || node.Children[c-'a'] == nil {
		return
	}

	n := node.Children[c-'a']

	if n.Word != "" {
		*result = append(*result, n.Word)
		n.Word = ""
	}

	if len(n.Children) > 0 {
		board[i][j] = '#'
		findWordsSearch(i+1, j, board, n, result)
		findWordsSearch(i-1, j, board, n, result)
		findWordsSearch(i, j+1, board, n, result)
		findWordsSearch(i, j-1, board, n, result)
		board[i][j] = c
	}
}

type FindWordsTrie struct {
	Children [26]*FindWordsTrie
	Word     string
}

/** Inserts a word into the trie. */
func (t *FindWordsTrie) Insert(word string) {
	node := t
	for _, l := range []byte(word) {
		l -= 'a'
		next := node.Children[l]

		if next == nil {
			next = new(FindWordsTrie)
			node.Children[l] = next
		}

		node = next
	}

	node.Word = word
}

// @lc code=end
