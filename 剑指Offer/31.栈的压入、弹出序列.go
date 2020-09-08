package offer

// 输入两个整数序列，第一个序列表示栈的压入顺序，请判断第二个序列是否为该栈的弹出顺序。假设压入栈的所有数字均不相等。例如，序列 {1,2,3,4,5} 是某栈的压栈序列，序列 {4,5,3,2,1} 是该压栈序列对应的一个弹出序列，但 {4,3,5,1,2} 就不可能是该压栈序列的弹出序列。

func validateStackSequences(pushed []int, popped []int) bool {
	if len(pushed) != len(popped) {
		return false
	}
	if len(pushed) == 0 {
		return true
	}
	stack := []int{}

	i, j := 0, 0
	for i < len(pushed) {
		if pushed[i] == popped[j] {
			i++
			j++

			for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
				stack = stack[:len(stack)-1]
				j++
			}
		} else {
			stack = append(stack, pushed[i])
			i++
		}
	}

	l := len(stack) - 1

	for k := 0; k < len(stack); k++ {
		if stack[l-k] != popped[k+j] {
			return false
		}
	}

	return true
}
