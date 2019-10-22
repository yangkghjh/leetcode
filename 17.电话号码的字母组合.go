package main

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) []string {
	codelist := []string{
		"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz",
	}

	result := []string{}
	for _, letter := range digits {
		num, _ := strconv.Atoi(string(letter))
		result = addLetters(result, codelist[num])
		// fmt.Println(num, codelist[num], result)
	}
	return result
}

func addLetters(in []string, codelist string) []string {
	var out []string
	if len(in) == 0 {
		out = make([]string, len(codelist))
		for i, code := range codelist {
			out[i] = string(code)
		}
		return out
	}

	out = make([]string, len(in)*len(codelist))
	var i int
	for _, str := range in {
		for _, code := range codelist {
			out[i] = str + string(code)
			i++
		}
	}

	return out
}

// @lc code=end
