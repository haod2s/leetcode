func combinationSum(candidates []int, target int) [][]int {
	var (
		dfs func(i, remaining int)
		ans [][]int
		t   []int
	)
	n := len(candidates)
	dfs = func(i, remaining int) {
		if remaining == 0 {
			ans = append(ans, append([]int(nil), t...))
			return
		}
		for j := i; j < n; j++ {
			if remaining-candidates[j] >= 0 {
				t = append(t, candidates[j])
				dfs(j, remaining-candidates[j])
				t = t[:len(t)-1]
			}
		}
	}
	dfs(0, target)

	return ans
}