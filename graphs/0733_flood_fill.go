package graphs

// 0733 - Flood Fill (Easy)
// Problem: An image is represented by an m x n integer grid image where image[i][j] represents the pixel value of the image.
// You are also given three integers sr, sc, and color. You should perform a flood fill on the image starting from the pixel image[sr][sc].
// To perform a flood fill:
// 1. Consider the starting pixel, plus any pixels connected 4-directionally to the starting pixel of the same color as the starting pixel,
// 2. plus any pixels connected 4-directionally to those pixels (also with the same color), and so on.
// 3. Replace the color of all of the aforementioned pixels with color.
// Return the modified image after performing the flood fill.

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if len(image) == 0 || len(image[0]) == 0 {
		return image
	}

	// Check if starting position is valid
	if sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[0]) {
		return image
	}

	originalColor := image[sr][sc]
	// If the original color is already the target color, no need to do anything
	if originalColor == color {
		return image
	}

	m, n := len(image), len(image[0])
	queue := [][2]int{{sr, sc}}
	image[sr][sc] = color

	// Directions: up, down, left, right
	directions := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		r, c := cell[0], cell[1]

		for _, dir := range directions {
			nr, nc := r+dir[0], c+dir[1]
			if nr >= 0 && nr < m && nc >= 0 && nc < n && image[nr][nc] == originalColor {
				image[nr][nc] = color
				queue = append(queue, [2]int{nr, nc})
			}
		}
	}

	return image
}

// Alternative DFS implementation
func floodFillDFS(image [][]int, sr int, sc int, color int) [][]int {
	if len(image) == 0 || len(image[0]) == 0 {
		return image
	}

	// Check if starting position is valid
	if sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[0]) {
		return image
	}

	originalColor := image[sr][sc]
	if originalColor == color {
		return image
	}

	dfs(image, sr, sc, originalColor, color)
	return image
}

func dfs(image [][]int, r, c, originalColor, newColor int) {
	if r < 0 || r >= len(image) || c < 0 || c >= len(image[0]) || image[r][c] != originalColor {
		return
	}

	image[r][c] = newColor

	// Recursively fill in all 4 directions
	dfs(image, r-1, c, originalColor, newColor) // up
	dfs(image, r+1, c, originalColor, newColor) // down
	dfs(image, r, c-1, originalColor, newColor) // left
	dfs(image, r, c+1, originalColor, newColor) // right
}