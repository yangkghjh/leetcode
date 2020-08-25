package leetcode

/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement1(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}

	return slow
}

// 时间复杂度： O(n)
// 空间复杂度： O(1)

// 因为 “元素的顺序可以改变”
// 当需要擦除的位置远小于元素总数时（总元素很多），可以使用将要去除的元素与数组尾部元素互换的
// 处理方式来减少元素的拷贝操作

func removeElement(nums []int, val int) int {
	n := len(nums)
	for i := 0; i < n; {
		if nums[i] == val {
			nums[i] = nums[n-1]
			n--
		} else {
			i++
		}
	}

	return n
}

// @lc code=end
