func diStringMatch(s string) []int {
	ans := make([]int, 0, len(s)+1)
	i, j := 0, len(s)
	for _, c := range s {
		if c == 'I' {
			ans = append(ans, i)
			i++
		} else {
			ans = append(ans, j)
			j--
		}
	}
	ans = append(ans, i)
	return ans
}