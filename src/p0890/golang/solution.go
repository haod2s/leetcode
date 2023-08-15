func findAndReplacePattern(words []string, pattern string) []string {
	ans := make([]string, 0)
	var (
		m       map[byte]byte
		matched bool
	)
	for _, w := range words {
		matched = true
		m = make(map[byte]byte)
		for i := 0; matched && i < len(w); i++ {
			if _, ok := m[w[i]]; !ok {
				m[w[i]] = pattern[i]
			} else if m[w[i]] != pattern[i] {
				matched = false
			}
		}
		m = make(map[byte]byte)
		for i := 0; matched && i < len(pattern); i++ {
			if _, ok := m[pattern[i]]; !ok {
				m[pattern[i]] = w[i]
			} else if m[pattern[i]] != w[i] {
				matched = false
			}
		}
		if matched {
			ans = append(ans, w)
		}
	}
	return ans
}