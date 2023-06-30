func closedIsland(grid [][]int) int {
	var (
		ans int
		dfs func(i, j int)
	)
	m, n := len(grid), len(grid[0])
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i == m || j == n || grid[i][j] == 1 {
			return
		}
		grid[i][j] = 1
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}
	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}
	for i := 0; i < n; i++ {
		dfs(0, i)
		dfs(m-1, i)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				ans++
				dfs(i, j)
			}
		}
	}

	return ans
}