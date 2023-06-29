func exist(board [][]byte, word string) bool {
	var (
		dfs func(i, j, idx int)
		ans bool
	)
	m, n := len(board), len(board[0])
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}
	dfs = func(i, j, idx int) {
		if idx == len(word) {
			ans = true
			return
		}
		if ans || i < 0 || j < 0 || i == m || j == n || vis[i][j] || board[i][j] != word[idx] {
			return
		}
		vis[i][j] = true
		dfs(i+1, j, idx+1)
		dfs(i-1, j, idx+1)
		dfs(i, j+1, idx+1)
		dfs(i, j-1, idx+1)
		vis[i][j] = false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, 0)
		}
	}

	return ans
}