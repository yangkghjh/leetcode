package leetcode

/*
 * @lc app=leetcode.cn id=752 lang=golang
 *
 * [752] 打开转盘锁
 */

// @lc code=start
func openLock(deadends []string, target string) int {
	hash := map[string]bool{}
	queue := []string{"0000", ""}

	for _, deadend := range deadends {
		hash[deadend] = true
	}

	// 操作可能性
	method := [][]int{{0, 1}, {1, 1}, {2, 1}, {3, 1},
		{0, -1}, {1, -1}, {2, -1}, {3, -1}}

	res := 0

	for i := 0; i < len(queue); i++ {
		if queue[i] == "" {
			if i == len(queue)-1 {
				break
			}
			queue = append(queue, "")
			res++
			continue
		}

		if queue[i] == target {
			return res
		}

		if _, ok := hash[queue[i]]; ok {
			continue
		}

		hash[queue[i]] = true

		for _, m := range method {
			d := openLockAdd(queue[i], m[0], m[1])
			if _, ok := hash[d]; !ok {
				queue = append(queue, d)
			}
		}
	}

	return -1
}

func openLockAdd(data string, pos, num int) string {
	d := []byte(data)
	if num == 1 {
		if d[pos] == '9' {
			d[pos] = '0'
		} else {
			d[pos]++
		}
	} else if num == -1 {
		if d[pos] == '0' {
			d[pos] = '9'
		} else {
			d[pos]--
		}
	}

	return string(d)
}

// @lc code=end
