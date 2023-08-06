func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	ans := 1
	for i, right := 1, points[0][1]; i < len(points); i++ {
		if right < points[i][0] {
			ans++
			right = points[i][1]
		}
	}
	return ans
}