package __19

//岛屿的最大面积 深度优先，广度优先
func maxAreaOfIsland(grid [][]int) int {
	max_area := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				max_area = max(max_area, dfs(grid, i, j))
			}
		}
	}
	return max_area
}

func dfs(grid [][]int, i int, j int) int {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == 0 {
		return 0
	}
	area := 1
	grid[i][j] = 0
	area += dfs(grid, i+1, j)
	area += dfs(grid, i, j+1)
	area += dfs(grid, i-1, j)
	area += dfs(grid, i, j-1)
	return area
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
