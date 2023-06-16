func combinationSum2(candidates []int, target int) [][]int {
	var (
		dfs func(i, remaining int)
		ans [][]int
		t   []int
	)
	dfs = func(i, remaining int) {
		if remaining == 0 {
			ans = append(ans, append([]int(nil), t...))
			return
		}
		for j := i; j < len(candidates); j++ {
			if j > i && candidates[j-1] == candidates[j] {
				continue
			}
			if remaining-candidates[j] >= 0 {
				t = append(t, candidates[j])
				dfs(j+1, remaining-candidates[j])
				t = t[:len(t)-1]
			}
		}
	}
	sort.Ints(candidates)
	dfs(0, target)

	return ans
}