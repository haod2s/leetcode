func matrixSum(nums [][]int) int {
	m, n := len(nums), len(nums[0])
	for i := 0; i < m; i++ {
		sort.Slice(nums[i], func(l, r int) bool {
			return nums[i][l] > nums[i][r]
		})
	}
	ans := 0
	for j := 0; j < n; j++ {
		maxItem := 0
		for i := 0; i < m; i++ {
			if nums[i][j] > maxItem {
				maxItem = nums[i][j]
			}
		}
		ans += maxItem
	}
	return ans
}