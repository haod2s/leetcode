func numIslands(grid [][]byte) int {
	var (
		cnt int
		dfs func(i, j int)
	)
	m, n := len(grid), len(grid[0])
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i == m || j == n || vis[i][j] || grid[i][j] == '0' {
			return
		}
		vis[i][j] = true
		for _, d := range directions {
			dfs(i+d[0], j+d[1])
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !vis[i][j] && grid[i][j] == '1' {
				cnt += 1
				dfs(i, j)
			}
		}
	}

	return cnt
}