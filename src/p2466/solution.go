package main

func solutionMemo(low int, high int, zero int, one int) int {
	const mod = 1e9 + 7

	memo := make(map[int]int)

	var dfs func(int) int
	dfs = func(l int) int {
		if _, ok := memo[l]; !ok {
			if l >= low && l <= high {
				memo[l] = (memo[l] + 1) % mod
			} else if l > high {
				return 0
			}

			memo[l] += (dfs(l+zero)%mod + dfs(l+one)%mod) % mod
		}

		return memo[l]
	}

	return dfs(0)
}
