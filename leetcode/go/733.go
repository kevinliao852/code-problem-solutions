func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	dfs(&image, sr, sc, color, image[sr][sc])

	return image
}

func dfs(image *[][]int, i, j, color, origin int) {
	fmt.Println(i, j)

	if i < 0 || i >= len(*image) || j < 0 || j >= len((*image)[0]) {
		return
	}

	if origin != (*image)[i][j] {
		return
	}

	if color == (*image)[i][j] {
		return
	}

	(*image)[i][j] = color

	// up
	dfs(image, i-1, j, color, origin)
	// right
	dfs(image, i, j+1, color, origin)
	// left
	dfs(image, i, j-1, color, origin)
	// down
	dfs(image, i+1, j, color, origin)
}
