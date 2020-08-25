package leetcode

/*
 * @lc app=leetcode.cn id=561 lang=golang
 *
 * [561] 数组拆分 I
 */

// @lc code=start
func arrayPairSum(nums []int) int {
	// nums 的范围远小于 nums 的长度时，可以使用额外的空间来
	// 索引数组（哈希表），之后分析哈希表来计算结果

	// 其他情况，先排序
	// sort.Ints(nums)
	quickSort(nums, 0, len(nums)-1)

	ans := 0
	for i := 0; i < len(nums); i += 2 {
		ans += nums[i]
	}

	return ans
}

func quickSort(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}

		if start < j {
			quickSort(arr, start, j)
		}
		if end > i {
			quickSort(arr, i, end)
		}
	}
}

// @lc code=end
