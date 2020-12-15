func minCostClimbingStairs(cost []int) int {
	m, n := cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		m, n = n, cost[i]+min(m, n)
	}
	return min(m, n)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}