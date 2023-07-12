func alternateDigitSum(n int) int {
	k := 0
	for x := n; x > 0; {
		k++
		x /= 10
	}
	sign := 1
	if k%2 == 0 {
		sign = -1
	}
	ans := 0
	for n > 0 {
		ans += sign * (n % 10)
		n /= 10
		sign *= -1
	}
	return ans
}