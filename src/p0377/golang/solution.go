package main

func solutionMemo(nums []int, target int) int {
	n := len(nums)

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, target+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
		memo[i][0] = 1
	}

	var dfs func(int, int) int
	dfs = func(i, remaining int) int {
		if remaining < 0 {
			return 0
		}
		if memo[i][remaining] == -1 && remaining > 0 {
			memo[i][remaining] = 0
			for j := 0; j < n; j++ {
				memo[i][remaining] += dfs(j, remaining-nums[j])
			}
		}

		return memo[i][remaining]
	}

	ans := dfs(0, target)

	return ans
}
