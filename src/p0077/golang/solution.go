func combine(n int, k int) [][]int {
	var (
		dfs func(i int)
		ans [][]int
		t   []int
	)
	dfs = func(i int) {
		if len(t) == k {
			ans = append(ans, append([]int(nil), t...))
			return
		}
		for j := i; j <= n; j++ {
			t = append(t, j)
			dfs(j + 1)
			t = t[:len(t)-1]
		}
	}
	dfs(1)

	return ans
}