package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	groupIndex := map[string][]string{}

	for _, str := range strs {
		sor := groupAnagramsSort(str)
		if g, ok := groupIndex[sor]; ok {
			groupIndex[sor] = append(g, str)
		} else {
			groupIndex[sor] = []string{str}
		}
	}

	result := make([][]string, len(groupIndex))
	i := 0
	for _, group := range groupIndex {
		result[i] = group
		i++
	}
	return result
}

func groupAnagramsSort(str string) string {
	s := groupAnagramslice(str)
	sort.Sort(s)
	return string(s)
}

type groupAnagramslice []byte

func (a groupAnagramslice) Len() int           { return len(a) }
func (a groupAnagramslice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a groupAnagramslice) Less(i, j int) bool { return a[i] < a[j] }

// @lc code=end
