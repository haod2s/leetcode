func equalSubstring(s string, t string, maxCost int) int {
	abs := func(x int) int {
		if x < 0 {
			return -1 * x
		}
		return x
	}
	n := len(s)
	ans := 0
	for i, j, sum := 0, 0, 0; j < n; j++ {
		sum += abs(int(s[j]) - int(t[j]))
		for sum > maxCost {
			sum -= abs(int(s[i]) - int(t[i]))
			i++
		}
		if j-i+1 > ans {
			ans = j - i + 1
		}
	}
	return ans
}