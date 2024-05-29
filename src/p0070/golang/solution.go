package main

func solutionMemo(n int) int {
	memo := make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int) int
	dfs = func(i int) int {
		if memo[i] == -1 {
			if i == 0 || i == 1 {
				memo[i] = 1
			} else {
				memo[i] = dfs(i-1) + dfs(i-2)
			}
		}

		return memo[i]
	}

	return dfs(n)
}

func solutionDP(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
