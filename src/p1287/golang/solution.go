func findSpecialInteger(arr []int) int {
	m := make(map[int]int)
	for _, v := range arr {
		if _, ok := m[v]; !ok {
			m[v] = 1
		} else {
			m[v] += 1
		}
	}
	var ret int
	for k, v := range m {
		if 4*v > len(arr) {
			ret = k
			break
		}
	}
	return ret
}