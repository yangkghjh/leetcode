package offer

func majorityElement(nums []int) int {
	x := 0
	votes := 0

	for _, num := range nums {
		if votes == 0 {
			x = num
		}

		if x == num {
			votes++
		} else {
			votes--
		}
	}

	return x
}
