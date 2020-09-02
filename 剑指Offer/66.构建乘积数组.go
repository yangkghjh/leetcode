package offer

// 给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，其中 B 中的元素 B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。不能使用除法。

func constructArr(a []int) []int {
	if len(a) == 0 {
		return []int{}
	}
	ans := make([]int, len(a))
	ans[0] = 1
	for i := 1; i < len(a); i++ {
		ans[i] = a[i-1] * ans[i-1]
	}

	tmp := 1
	for i := len(a) - 2; i >= 0; i-- {
		tmp *= a[i+1]
		ans[i] = ans[i] * tmp
	}

	return ans
}
