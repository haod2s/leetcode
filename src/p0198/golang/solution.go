package main

func rob(nums []int) int {
	n := len(nums)
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][1] = nums[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = max(dp[i-1][0]+nums[i], dp[i-1][1])
	}

	return max(dp[n-1][0], dp[n-1][1])
}

func solutionMemo(nums []int) int {
	n := len(nums)

	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int) int
	dfs = func(i int) int {
		if memo[i] == -1 {
			if i == 0 {
				memo[i] = nums[0]
			} else if i == 1 {
				memo[i] = max(nums[0], nums[1])
			} else {
				memo[i] = max(dfs(i-2)+nums[i], dfs(i-1))
			}
		}

		return memo[i]
	}

	return dfs(n - 1)
}

func solutionDP(nums []int) int {
	n := len(nums)

	prev, cur := -1, nums[0]
	if n > 1 {
		prev, cur = nums[0], max(nums[0], nums[1])
	}

	for i := 2; i < n; i++ {
		prev, cur = cur, max(prev+nums[i], cur)
	}

	return cur
}
