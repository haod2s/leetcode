func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	var dfs func(i, j, path int)
	minPath := math.MaxInt
	mem := make([][]int, n)
	for i := 0; i < n; i++ {
		mem[i] = make([]int, n)
		for j := 0; j < n; j++ {
			mem[i][j] = math.MaxInt
		}
	}
	dfs = func(i, j, path int) {
		if i == n {
			if path < minPath {
				minPath = path
			}
			return
		}
		if j < 0 || j == n || path >= mem[i][j] {
			return
		}
		mem[i][j] = path
		dfs(i+1, j-1, path+matrix[i][j])
		dfs(i+1, j, path+matrix[i][j])
		dfs(i+1, j+1, path+matrix[i][j])
	}
	for j := 0; j < n; j++ {
		dfs(0, j, 0)
	}
	return minPath
}