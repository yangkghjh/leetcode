package offer

func lastRemaining(n int, m int) int {
	if n == 1 {
		return 0
	}
	return (lastRemaining(n-1, m) + m) % n
}
