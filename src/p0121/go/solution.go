func maxProfit(prices []int) int {
	min, ans := prices[0], 0
	for _, v := range prices {
		if min > v {
			min = v
		} else if v-min > ans {
			ans = v - min
		}
	}

	return ans
}