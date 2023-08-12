func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	idx := make([]int, len(indices))
	for i := 0; i < len(indices); i++ {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return indices[idx[i]] > indices[idx[j]]
	})
	ans := []byte(s)
	for _, i := range idx {
		j := indices[i]
		if j+len(sources[i])-1 < len(ans) && string(ans[j:j+len(sources[i])]) == sources[i] {
			before := ans[0:j]
			after := ans[j+len(sources[i]):]
			ans = append(before, append([]byte(targets[i]), after...)...)
		}
	}
	return string(ans)
}