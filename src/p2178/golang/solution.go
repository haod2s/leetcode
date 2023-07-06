func maximumEvenSplit(finalSum int64) []int64 {
	ans := make([]int64, 0)
	if finalSum%2 != 0 {
		return ans
	}
	for i := int64(2); finalSum-i >= 0; i += 2 {
		ans = append(ans, i)
		finalSum -= i
	}
	ans[len(ans)-1] += finalSum
	return ans
}