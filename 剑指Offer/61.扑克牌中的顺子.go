package offer

// 从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。

func isStraight(nums []int) bool {
	hash := make([]bool, 14)

	max, min := 0, 14
	zero := false

	for _, n := range nums {
		if n == 0 {
			zero = true
			continue
		}
		if hash[n] {
			return false
		}

		hash[n] = true

		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	if zero {
		return max-min <= 4
	}

	return max-min == 4
}
