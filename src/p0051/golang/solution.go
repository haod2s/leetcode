func solveNQueens(n int) [][]string {
	var (
		dfs func(int)
		pos []int
		ans [][]string
	)
	toBoard := func() []string {
		board := make([]string, n)
		for i := 0; i < n; i++ {
			board[i] = strings.Repeat(string('.'), n)
			b := []byte(board[i])
			b[pos[i]] = byte('Q')
			board[i] = string(b)
		}
		return board
	}
	cols := make([]bool, n)
	main := make([]bool, 2*n-1)
	sub := make([]bool, 2*n-1)
	dfs = func(row int) {
		if row == n {
			ans = append(ans, toBoard())
			return
		}
		for i := 0; i < n; i++ {
			if !cols[i] && !main[row-i+n-1] && !sub[row+i] {
				cols[i], main[row-i+n-1], sub[row+i] = true, true, true
				pos = append(pos, i)
				dfs(row + 1)
				pos = pos[:len(pos)-1]
				cols[i], main[row-i+n-1], sub[row+i] = false, false, false
			}
		}
	}
	dfs(0)

	return ans
}