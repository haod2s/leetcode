func minimumOperations(num string) int {
	if num == "0" {
		return 0
	}
	n := len(num)
	ans := n
	if strings.Contains(num, "0") {
		ans = n - 1
	}
	for i := n - 1; i >= 0; i-- {
		if num[i] == '0' {
			for j := i - 1; j >= 0; j-- {
				if num[j] == '0' {
					ans = min(ans, n-j-2)
				}
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		if num[i] == '5' {
			for j := i - 1; j >= 0; j-- {
				if num[j] == '2' {
					ans = min(ans, n-j-2)
				}
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		if num[i] == '0' {
			for j := i - 1; j >= 0; j-- {
				if num[j] == '5' {
					ans = min(ans, n-j-2)
				}
			}
		}
	}
	for i := n - 1; i >= 0; i-- {
		if num[i] == '5' {
			for j := i - 1; j >= 0; j-- {
				if num[j] == '7' {
					ans = min(ans, n-j-2)
				}
			}
		}
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}