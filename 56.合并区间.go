package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	s := interval(intervals)
	sort.Sort(s)
	intervals = s
	index := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= intervals[index][1] {
			if intervals[i][1] > intervals[index][1] {
				intervals[index][1] = intervals[i][1]
			}
		} else {
			index++
			intervals[index] = intervals[i]
		}
	}

	return intervals[0 : index+1]
}

type interval [][]int

func (a interval) Len() int           { return len(a) }
func (a interval) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a interval) Less(i, j int) bool { return a[i][0] < a[j][0] }

// @lc code=end
