func largestTimeFromDigits(arr []int) string {
	var (
		ans [][]int
		dfs func(i int)
		t   []int
	)
	sort.Ints(arr)
	vis := make([]bool, 4)
	dfs = func(i int) {
		if i == 4 {
			if t[0] > 2 || (t[0] == 2 && t[1] > 3) || t[2] > 5 {
				return
			}
			ans = append(ans, append([]int(nil), t...))
			return
		}
		for j := 0; j < 4; j++ {
			if vis[j] || j > 0 && !vis[j-1] && arr[j] == arr[j-1] {
				continue
			}
			t = append(t, arr[j])
			vis[j] = true
			dfs(i + 1)
			vis[j] = false
			t = t[:len(t)-1]
		}
	}
	dfs(0)
	sort.Slice(ans, func(i, j int) bool {
		for k := 0; k < 4; k++ {
			if ans[i][k] == ans[j][k] {
				continue
			}
			return ans[i][k] < ans[j][k]
		}
		return false
	})
	if len(ans) == 0 {
		return ""
	}
	x := ans[len(ans)-1]
	return fmt.Sprintf("%c%c:%c%c", '0'+byte(x[0]), '0'+byte(x[1]), '0'+byte(x[2]), '0'+byte(x[3]))
}