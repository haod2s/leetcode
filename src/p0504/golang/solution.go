func convertToBase7(num int) string {
	sign := 1
	ans := make([]byte, 0)
	if num < 0 {
		sign = -1
		num *= sign
	}
	for num >= 7 {
		ans = append(ans, '0'+byte(num%7))
		num /= 7
	}
	ans = append(ans, '0'+byte(num))
	for i, j := 0, len(ans)-1; i < j; {
		ans[i], ans[j] = ans[j], ans[i]
		i++
		j--
	}
	if sign < 0 {
		return "-" + string(ans)
	}
	return string(ans)
}

// 腾讯