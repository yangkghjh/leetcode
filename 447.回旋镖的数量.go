package leetcode

/*
 * @lc app=leetcode.cn id=447 lang=golang
 *
 * [447] 回旋镖的数量
 */

// @lc code=start
func numberOfBoomerangs(points [][]int) int {
	hash := map[int]int{}
	res := 0
	gendistance := func(a, b []int) int {
		return (a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1])
	}

	for i := 0; i < len(points); i++ {
		hash = map[int]int{}
		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}
			distance := gendistance(points[i], points[j])
			_, ok := hash[distance]

			if ok {
				res += hash[distance] * 2
				hash[distance]++
			} else {
				hash[distance] = 1
			}
		}
	}

	return res
}

// @lc code=end
