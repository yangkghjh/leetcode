package offer

// 把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。
// 你需要用一个浮点数数组返回答案，其中第 i 个元素代表这 n 个骰子所能掷出的点数集合中第 i 小的那个的概率。

func twoSum(n int) []float64 {
	t := n
	m := map[int]int{0: 1}
	for t > 0 {
		m1 := map[int]int{}
		for i := 1; i <= 6; i++ {
			for k, v := range m {
				m1[k+i] += v
			}
		}
		m = m1
		t--
	}
	sum := 0
	for _, v := range m {
		sum += v
	}
	ret := make([]float64, n*6-n+1)
	for k, v := range m {
		ret[k-n] = float64(v) / float64(sum)
	}
	return ret
}
