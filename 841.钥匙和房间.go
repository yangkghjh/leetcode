package leetcode

/*
 * @lc app=leetcode.cn id=841 lang=golang
 *
 * [841] 钥匙和房间
 */

// @lc code=start
func canVisitAllRooms(rooms [][]int) bool {
	nums := len(rooms)
	if nums <= 1 {
		return true
	}

	visited := map[int]bool{}

	queue := []int{0}

	for i := 0; i < len(queue) && len(visited) < nums; i++ {
		n := queue[i]
		visited[n] = true
		for _, key := range rooms[n] {
			if _, ok := visited[key]; !ok && key < nums {
				queue = append(queue, key)
			}
		}
	}

	return len(visited) == nums
}

// 时间复杂度 低于 O(钥匙总数+房间总数)
// 空间复杂度 低于 O(房间总数)，如果可以修改入参，可以省去 visited

// @lc code=end
