package offer

import (
	"sort"
	"strconv"
)

type minNumberInts []int

func (m minNumberInts) Len() int      { return len(m) }
func (m minNumberInts) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

func (m minNumberInts) Less(i, j int) bool {
	x := strconv.Itoa(m[i])
	y := strconv.Itoa(m[j])
	xy, _ := strconv.Atoi(x + y)
	yx, _ := strconv.Atoi(y + x)

	return xy < yx
}

func minNumber(nums []int) string {
	sort.Sort(minNumberInts(nums))

	ans := ""
	for _, num := range nums {
		ans += strconv.Itoa(num)
	}

	return ans
}
