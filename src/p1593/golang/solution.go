func maxUniqueSplit(s string) int {
	var (
		ans int
		dfs func(idx, i int)
	)
	vis := map[string]struct{}{}
	dfs = func(idx, i int) {
		if i == len(s) {
			if idx > ans {
				ans = idx
			}
			return
		}
		for j := i + 1; j <= len(s); j++ {
			w := string(s[i:j])
			if _, ok := vis[w]; !ok {
				vis[w] = struct{}{}
				dfs(idx+1, j)
				delete(vis, w)
			}
		}
	}
	dfs(0, 0)

	return ans
}