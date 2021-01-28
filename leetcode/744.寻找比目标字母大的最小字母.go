package leetcode

/*
 * @lc app=leetcode.cn id=744 lang=golang
 *
 * [744] 寻找比目标字母大的最小字母
 */

// @lc code=start
func nextGreatestLetter(letters []byte, target byte) byte {
	if len(letters) == 1 {
		return letters[0]
	}
	left, right := 0, len(letters)-2

	for left <= right {
		mid := (left + right) / 2

		if letters[mid] <= target && letters[mid+1] > target {
			return letters[mid+1]
		}

		if letters[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return letters[0]
}

// @lc code=end
