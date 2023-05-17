func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		if ans[len(ans)-1][1] >= intervals[i][1] {
			continue
		}
		if ans[len(ans)-1][1] >= intervals[i][0] {
			ans[len(ans)-1][1] = intervals[i][1]
		} else if ans[len(ans)-1][1] < intervals[i][0] {
			ans = append(ans, intervals[i])
		}
	}

	return ans
}

// 华为