func maxProfit(prices []int, fee int) int {
	N := len(prices)
	dp := make([][]int, N)
	dp[0] = make([]int, 2)
	dp[0][0], dp[0][1] = 0, -1*prices[0]
	for i := 1; i < N; i++ {
		dp[i] = make([]int, 2)
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[N-1][0]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}