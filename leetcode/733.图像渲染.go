package leetcode

/*
 * @lc app=leetcode.cn id=733 lang=golang
 *
 * [733] 图像渲染
 */

// @lc code=start
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	// len(image) [1,50]
	// len(image[0]) [1,50]
	queue := [][]int{{sr, sc}}
	color := image[sr][sc]
	if color == newColor {
		return image
	}
	image[sr][sc] = newColor
	width := len(image[0])
	height := len(image)

	for i := 0; i < len(queue); i++ {
		r, c := queue[i][0], queue[i][1]
		if r >= 1 && image[r-1][c] == color {
			image[r-1][c] = newColor
			queue = append(queue, []int{r - 1, c})
		}
		if r < height-1 && image[r+1][c] == color {
			image[r+1][c] = newColor
			queue = append(queue, []int{r + 1, c})
		}
		if c >= 1 && image[r][c-1] == color {
			image[r][c-1] = newColor
			queue = append(queue, []int{r, c - 1})
		}
		if c < width-1 && image[r][c+1] == color {
			image[r][c+1] = newColor
			queue = append(queue, []int{r, c + 1})
		}
	}

	return image
}

// @lc code=end
