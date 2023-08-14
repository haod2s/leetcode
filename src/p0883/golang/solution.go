func projectionArea(grid [][]int) int {
	n := len(grid)
	area := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				area++
			}
		}
	}
	for i := 0; i < n; i++ {
		hz := 0
		for j := 0; j < n; j++ {
			if grid[i][j] > hz {
				hz = grid[i][j]
			}
		}
		area += hz
	}
	for i := 0; i < n; i++ {
		hz := 0
		for j := 0; j < n; j++ {
			if grid[j][i] > hz {
				hz = grid[j][i]
			}
		}
		area += hz
	}
	return area
}