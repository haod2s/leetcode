func maxLengthBetweenEqualCharacters(s string) int {
	m := make(map[rune][]int)
	for i, c := range s {
		if _, ok := m[c]; !ok {
			m[c] = []int{-1, -1}
			m[c][0] = i
		} else {
			m[c][1] = i
		}
	}
	ans := -1
	for _, v := range m {
		if v[1] != -1 && v[1]-v[0]-1 > ans {
			ans = v[1] - v[0] - 1
		}
	}
	return ans
}