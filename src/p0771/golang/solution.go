func numJewelsInStones(jewels string, stones string) int {
	m := map[rune]struct{}{}
	for _, v := range jewels {
		m[v] = struct{}{}
	}
	ans := 0
	for _, v := range stones {
		if _, ok := m[v]; ok {
			ans += 1
		}
	}
	return ans
}