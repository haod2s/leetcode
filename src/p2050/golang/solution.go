func minimumTime(n int, relations [][]int, time []int) int {
	edges := make([][]int, n+1)
	// indegree
	in := make([]int, n+1)
	for _, r := range relations {
		edges[r[0]] = append(edges[r[0]], r[1])
		in[r[1]]++
	}
	dp := make([]int, n+1)
	q := make([]int, 0)
	for i := 1; i <= n; i++ {
		if in[i] == 0 {
			q = append(q, i)
			dp[i] = time[i-1]
		}
	}
	ans := math.MinInt
	for len(q) > 0 {
		i := q[0]
		q = q[1:]
		ans = max(ans, dp[i])
		for _, v := range edges[i] {
			in[v]--
			dp[v] = max(dp[v], dp[i]+time[v-1])
			if in[v] == 0 {
				q = append(q, v)
			}
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}