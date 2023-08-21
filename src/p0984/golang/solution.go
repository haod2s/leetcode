func strWithout3a3b(a int, b int) string {
	ans := make([]byte, 0, a+b)
	for a > b && b > 0 {
		ans = append(ans, []byte("aab")...)
		a -= 2
		b--
	}
	for b > a && a > 0 {
		ans = append(ans, []byte("bba")...)
		b -= 2
		a--
	}
	for a > 0 && b > 0 {
		ans = append(ans, []byte("ab")...)
		a--
		b--
	}
	for ; a > 0; a-- {
		ans = append(ans, 'a')
	}
	for ; b > 0; b-- {
		ans = append(ans, 'b')
	}
	return string(ans)
}