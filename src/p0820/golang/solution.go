func minimumLengthEncoding(words []string) int {
	m := make(map[string]int)
	temp := make([]string, 0)
	for _, w := range words {
		if _, ok := m[w]; !ok {
			temp = append(temp, w)
			m[w] = len(w)
		}
	}
	sort.Slice(temp, func(i, j int) bool {
		return len(temp[i]) > len(temp[j])
	})
	for i := 0; i < len(temp); i++ {
		if _, ok := m[temp[i]]; ok {
			for j := i + 1; j < len(temp); j++ {
				if strings.HasSuffix(temp[i], temp[j]) {
					delete(m, temp[j])
				}
			}
		}
	}
	ans := 0
	for _, v := range m {
		ans += v + 1
	}

	return ans
}

// 腾讯