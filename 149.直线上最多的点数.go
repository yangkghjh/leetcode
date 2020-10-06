package leetcode

/*
 * @lc app=leetcode.cn id=149 lang=golang
 *
 * [149] 直线上最多的点数
 */

// @lc code=start
func maxPoints(points [][]int) int {
	if len(points) < 2 {
		return len(points)
	}

	res := 1
	hash := map[K]int{}

	samePoint := K{0, 0}
	slope := func(p1, p2 []int) K {
		dx := int64(p1[0] - p2[0])
		dy := int64(p1[1] - p2[1])
		if dx == 0 && dy == 0 {
			return K{int64(0), int64(0)}
		} //p1,p2是同一点
		if dx == 0 {
			return K{int64(1), int64(0)}
		} //p1,p2斜率为无穷
		if dy == 0 {
			return K{int64(0), int64(1)}
		} //p1,p2斜率为0
		if dx < 0 { //统一设置dx为正
			dx = -dx
			dy = -dy
		}
		d := gcd(dx, dy)
		return K{dx / d, dy / d} //返回最简约之比
	}

	for i := 0; i < len(points)-1; i++ {
		hash = map[K]int{}
		max := 1
		samepoints := 0
		for j := i + 1; j < len(points); j++ {
			s := slope(points[i], points[j])

			// fmt.Println(i, j, s, max, s == samePoint)
			if s == samePoint {
				samepoints++
				continue
			}

			if _, ok := hash[s]; ok {
				hash[s]++
			} else {
				hash[s] = 2
			}

			if hash[s] > max {
				max = hash[s]
			}
		}

		max += samepoints

		if max > res {
			res = max
		}
		// fmt.Println(res)
	}

	return res
}

//最大公约数
func gcd(a, b int64) int64 {
	if a == 0 {
		return b
	}
	return gcd(b%a, a)
}

//斜率，用最简约分子分母表示，为防止溢出，类型为int64
type K struct {
	m int64 //分子
	n int64 //分母
}

// @lc code=end
