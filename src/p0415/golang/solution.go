func addStrings(num1 string, num2 string) string {
	ans := make([]byte, 0)
	flag := 0
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || flag == 1; i, j = i-1, j-1 {
		x := flag
		if i >= 0 {
			x += int(num1[i] - '0')
		}
		if j >= 0 {
			x += int(num2[j] - '0')
		}
		flag = x / 10
		ans = append(ans, byte(x%10)+'0')
	}
	for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}
	return string(ans)
}