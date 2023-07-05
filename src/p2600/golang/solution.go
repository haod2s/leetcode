func kItemsWithMaximumSum(numOnes int, numZeros int, numNegOnes int, k int) int {
	ans := 0
	for k > 0 {
		if numOnes > 0 {
			ans += 1
			numOnes--
		} else if numZeros > 0 {
			numZeros--
		} else {
			ans -= 1
			numNegOnes--
		}
		k--
	}
	return ans
}