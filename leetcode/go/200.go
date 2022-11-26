func numIslands(grid [][]byte) int {
	var count int
	visited := make([][]bool, len(grid))

	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid[0]))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if visited[i][j] == false && grid[i][j] != 48 {
				dfs(&grid, i, j, &visited)
				count++
			}
		}
	}

	return count
}

func dfs(grid *[][]byte, i, j int, visited *[][]bool) {
	if i < 0 || j < 0 || i >= len(*grid) || j >= len((*grid)[0]) {
		return
	}

	if (*grid)[i][j] == 48 {
		return
	}

	if (*visited)[i][j] == true {
		return
	}

	(*visited)[i][j] = true

	dfs(grid, i-1, j, visited)
	dfs(grid, i, j+1, visited)
	dfs(grid, i+1, j, visited)
	dfs(grid, i, j-1, visited)

}
