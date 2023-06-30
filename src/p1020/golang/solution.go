func numEnclaves(grid [][]int) int {
	var dfs func(i, j int)
	m, n := len(grid), len(grid[0])
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i == m || j == n || grid[i][j] == 0 {
			return
		}
		grid[i][j] = 0
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}
	for i := 0; i < m; i++ {
		if grid[i][0] == 1 {
			dfs(i, 0)
		}
		if grid[i][n-1] == 1 {
			dfs(i, n-1)
		}
	}
	for i := 0; i < n; i++ {
		if grid[0][i] == 1 {
			dfs(0, i)
		}
		if grid[m-1][i] == 1 {
			dfs(m-1, i)
		}
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans += grid[i][j]
		}
	}

	return ans
}