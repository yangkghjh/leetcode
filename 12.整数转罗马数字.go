package leetcode
/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */
func intToRoman(num int) string {
	if num > 3999 || num < 1 {
		return ""
	}
	var roman string
	kiloBit := num / 1000
	for i := 0; i < kiloBit; i++ {
		roman += "M"
	}

	num = num % 1000
	hundredBit := num / 100
	roman += transformBit(hundredBit, 'C', 'D', 'M')

	num = num % 100
	dacadeBit := num / 10
	roman += transformBit(dacadeBit , 'X', 'L', 'C')

	lastBit := num % 10
	roman += transformBit(lastBit, 'I', 'V', 'X')

	return roman
}

func transformBit(bit int, baseCode, halfCode, entireCode byte) string {
	var roman []byte 
	if bit == 0 {
		// do nothing
	} else if bit == 4 {
		roman = append(roman, baseCode, halfCode) 
	} else if bit == 9 {
		roman = append(roman, baseCode, entireCode) 
	} else if bit == 5 {
		roman = append(roman, halfCode) 
	} else {
		if bit >= 6 {
			roman = append(roman, halfCode)
			bit -= 5
		}
		for i := 0; i < bit; i++ {
			roman = append(roman, baseCode) 
		}
	}

	return string(roman)
}

