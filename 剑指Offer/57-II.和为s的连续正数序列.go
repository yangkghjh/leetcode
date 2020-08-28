package offer

// 输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

func findContinuousSequence(target int) [][]int {
	if target == 1 {
		return [][]int{[]int{1}}
	}
	start, end, sum := 1, 2, 3

	ans := [][]int{}

	for start < end && end <= target/2+1 {
		// fmt.Println(start, end)
		if sum == target {
			ans = append(ans, getList(start, end))
			end++
			sum += end
		} else if sum < target {
			end++
			sum += end
		} else {
			sum -= start
			start++
		}
	}

	return ans
}

func getList(start, end int) []int {
	l := end - start + 1
	ans := make([]int, l)

	for i := 0; i < l; i++ {
		ans[i] = i + start
	}

	return ans
}
