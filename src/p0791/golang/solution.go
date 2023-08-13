func customSortString(order string, s string) string {
	idx := make([]int, 26)
	for i := 0; i < len(order); i++ {
		idx[order[i]-'a'] = i
	}
	tmp := []byte(s)
	sort.Slice(tmp, func(i, j int) bool {
		return idx[tmp[i]-'a'] < idx[tmp[j]-'a']
	})
	return string(tmp)
}