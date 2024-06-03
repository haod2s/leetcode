package main

func solutionDP(pressedKeys string) int {
	const (
		mod = 1e9 + 7
		mx  = 1e5
	)

	dp1 := [mx + 1]int{1, 1, 2, 4}
	dp2 := [mx + 1]int{1, 1, 2, 4}

	for i := 4; i <= mx; i++ {
		dp1[i] = (dp1[i-1] + dp1[i-2] + dp1[i-3] + dp1[i-4]) % mod
		dp2[i] = (dp2[i-1] + dp2[i-2] + dp2[i-3]) % mod
	}

	n := len(pressedKeys)
	ans := 1
	for i := 0; i < n; {
		j := i + 1
		for ; j < n && pressedKeys[j-1] == pressedKeys[j]; j++ {
		}

		if pressedKeys[i] == '7' || pressedKeys[i] == '9' {
			ans = ans * dp1[j-i] % mod
		} else {
			ans = ans * dp2[j-i] % mod
		}

		i = j
	}

	return ans
}
