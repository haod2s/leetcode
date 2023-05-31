func longestPalindrome(s string) int {
	m := make(map[byte]int)
	for _, c := range s {
		m[byte(c)] += 1
	}
	var (
		largestKey byte
		largestOdd int
		ans        int
	)
	for k, v := range m {
		if v%2 == 1 && v > largestOdd {
			largestOdd = v
			largestKey = k
		}
	}
	ans = largestOdd
	delete(m, largestKey)
	for _, v := range m {
		if v%2 == 0 {
			ans += v
		} else if largestOdd != 1 && v != 1 {
			ans += v - 1
		}
	}

	return ans
}

// 腾讯