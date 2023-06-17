func subsets(nums []int) [][]int {
	var (
		dfs func(i int)
		ans [][]int
		t   []int
	)
	dfs = func(i int) {
		ans = append(ans, append([]int(nil), t...))
		for j := i; j < len(nums); j++ {
			t = append(t, nums[j])
			dfs(j + 1)
			t = t[:len(t)-1]
		}
	}
	dfs(0)

	return ans
}