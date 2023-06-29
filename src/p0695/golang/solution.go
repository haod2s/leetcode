func maxAreaOfIsland(grid [][]int) int {
	var (
		ans int
		dfs func(i, j int) int
	)
	m, n := len(grid), len(grid[0])
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}
	directtions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	dfs = func(i, j int) int {
		area := 0
		if i < 0 || j < 0 || i == m || j == n || vis[i][j] || grid[i][j] == 0 {
			return area
		}
		area += 1
		vis[i][j] = true
		for _, d := range directtions {
			area += dfs(i+d[0], j+d[1])
		}
		return area
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if area := dfs(i, j); area > ans {
				ans = area
			}
		}
	}

	return ans
}