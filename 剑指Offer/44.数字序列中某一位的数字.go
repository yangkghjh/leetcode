package offer

func findNthDigit(n int) int {
	// 0 1
	// 1-9 9
	// 10-99 90
	// 100-999 900

	k := 9
	l := 1
	for n > k*l {
		n -= k * l
		k *= 10
		l++
	}

	num := (n-1)/l + k/9
	pos := (n - 1) % l

	for pos < l-1 {
		num /= 10
		pos++
	}

	return num % 10
}
