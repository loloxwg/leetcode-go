package __19

/*200. 岛屿数量*/
func numIslands(grid [][]byte) int {
	counter := 0
	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0]) && grid[i][j] == '1' {
			grid[i][j] = '0'
			dfs(i, j+1)
			dfs(i, j-1)
			dfs(i+1, j)
			dfs(i-1, j)
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				counter++
				dfs(i, j)
			}
		}
	}

	return counter

}
