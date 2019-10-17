package main

/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	if s == "" {
		return 0
	}

	var result int
	romans := []byte(s)
	var i int
	for ; i < len(romans); i++ {
		if romans[i] == 'M' {
			result += 1000
		} else {
			break
		}
	}

	romans = romans[i:]
	for _, key := range []int{100, 10, 1} {
		value := romansList[key]
		if len(romans) == 0 {
			return result
		}

		var byteResult int
		romans, byteResult = transformRoman(romans, key, value[0], value[1], value[2])
		// fmt.Println(string(romans), byteResult)
		result += byteResult
	}

	return result
}

var romansList = map[int][]byte{
	1:   []byte{'I', 'V', 'X'},
	10:  []byte{'X', 'L', 'C'},
	100: []byte{'C', 'D', 'M'},
}

func transformRoman(romans []byte, value int, baseCode, halfCode, entireCode byte) ([]byte, int) {
	var i int
	var result int
	var hasBaseCode bool
	for ; i < len(romans); i++ {
		if romans[i] == baseCode {
			result += value
			hasBaseCode = true
		} else if romans[i] == halfCode {
			if hasBaseCode {
				result += value * 3
			} else {
				result += value * 5
			}
		} else if romans[i] == entireCode {
			if hasBaseCode {
				result += value * 8
			}
		} else {
			break
		}
	}

	return romans[i:], result
}

// @lc code=end
